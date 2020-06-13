package app
import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

const configHeader = `# %s configuration
# 
# This config has https://github.com/tpphu/gobox syntax.
# Empty lines or lines starting with # will be ignored.
# All other lines must look like KEY=VALUE.
# The VALUE must not be enclosed in quotes as well!
`

var (
	openOrCreate = os.OpenFile
)

type AppFlagSet struct {
	*flag.FlagSet
}

func newFlagSet(name string, fs *flag.FlagSet) *AppFlagSet {
	fSet := &AppFlagSet{fs}
	//fSet.Usage = flag_customUsage(name, fSet)
	return fSet
}

func saveConfig(w io.Writer, obsKeys map[string]string) {
	// find flags pointing to the same variable. We will only write the longest
	// named flag to the config file, the shorthand version is ignored.
	deduped := make(map[flag.Value]flag.Flag)
	flag.VisitAll(func(f *flag.Flag) {
		if cur, ok := deduped[f.Value]; !ok || utf8.RuneCountInString(f.Name) > utf8.RuneCountInString(cur.Name) {
			deduped[f.Value] = *f
		}
	})
	flag.VisitAll(func(f *flag.Flag) {
		if cur, ok := deduped[f.Value]; ok && cur.Name == f.Name {
			_, usage := flag.UnquoteUsage(f)
			usage = strings.Replace(usage, "\n    \t", "\n# ", -1)
			fmt.Fprintf(w, "\n## %s (default %v)\n", usage, f.DefValue)
			fmt.Fprintf(w, "%s=%v\n", f.Name, f.Value.String())
		}
	})

	// if we have obsolete keys left from the old config, preserve them in an
	// additional section at the end of the file
	if obsKeys != nil && len(obsKeys) > 0 {
		fmt.Fprintln(w, "\n\n# The following options are probably deprecated and not used currently!")
		for key, val := range obsKeys {
			fmt.Fprintf(w, "#%v=%v\n", key, val)
		}
	}
}

func parseConfig(r io.Reader) map[string]string {
	obsKeys := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}

		// find first assignment symbol and parse key, val
		i := strings.IndexAny(line, "=:")
		if i == -1 {
			continue
		}
		key, val := strings.TrimSpace(line[:i]), strings.TrimSpace(line[i+1:])

		if err := flag.Set(key, val); err != nil {
			obsKeys[key] = val
			continue
		}
	}
	return obsKeys
}

func getPathDefault() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fileName := ""
	if len(os.Args) > 1 {
		args := os.Args[1:]
		if len(args) == 1 {
			path += "/.env"
			return path
		}
		for key, arg := range args {
			if arg == "outenv" {
				fileName = args[key+1]
			}
		}
	}
	path += "/" + fileName
	return path
}

func (f *AppFlagSet) Parse(appName string) error {
	if f.Parsed() {
		return fmt.Errorf("flags have been parsed already")
	}
	cPath := getPathDefault()
	cf, err := openOrCreate(cPath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("unable to open %s config file %v for reading and writing: %v", appName, cPath, err)
	}
	defer cf.Close()
	// read config to buffer and parse
	oldConf := new(bytes.Buffer)
	obsoleteKeys := parseConfig(io.TeeReader(cf, oldConf))
	// write updated config to another buffer
	newConf := new(bytes.Buffer)
	fmt.Fprintf(newConf, configHeader, appName)
	saveConfig(newConf, obsoleteKeys)
	// only write the file if it changed
	if !bytes.Equal(oldConf.Bytes(), newConf.Bytes()) {
		if ofs, err := cf.Seek(0, 0); err != nil || ofs != 0 {
			return fmt.Errorf("failed to seek to beginning of %s: %v", cPath, err)
		} else if err = cf.Truncate(0); err != nil {
			return fmt.Errorf("failed to truncate %s: %v", cPath, err)
		} else if _, err = newConf.WriteTo(cf); err != nil {
			return fmt.Errorf("failed to write %s: %v", cPath, err)
		}
	}
	flag.Parse()
	return nil
}
