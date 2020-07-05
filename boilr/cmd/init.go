package cmd

import (
	"fmt"
	cli "github.com/spf13/cobra"
	"github.com/tpphu/gobox/boilr/util/exit"
	"github.com/tpphu/gobox/boilr/util/log"
	osutil "github.com/tpphu/gobox/boilr/util/os"
	"github.com/tpphu/gobox/boilr/util/validate"
	"path/filepath"
)

var Init = &cli.Command{
	Use:   "init",
	Short: "Initialize template required by gobox",
	Run: func(c *cli.Command, _ []string) {
		serviceName := GetStringFlag(c, "service-name")
		resource := GetStringFlag(c, "resource")
		//Validate args
		MustValidateArgs([]string{serviceName, resource}, []validate.Argument{
			{"service-name", validate.AlphanumericExt},
			{"resource", validate.AlphanumericOpt1},
		})
		path := GetStringFlag(c, "path")
		if path == "" {
			log.Info("path is empty. It will be init in current directory")
			path, _ = osutil.Getwd()
		}
		if exists, err := osutil.DirExists(path); exists {
			if shouldRecreate := GetBoolFlag(c, "force"); !shouldRecreate {
				exit.GoodEnough("template registry is already initialized use -f to reinitialize")
			}
		} else if err != nil {
			exit.Error(fmt.Errorf("init: %s", err))
		}
		//Get template path
		tempPath, err := getTemplatePath()
		if err != nil {
			exit.Error(fmt.Errorf("getTemplatePath: %s", err))
		}
		//End
		//Join path and service name
		svPath := filepath.Join(path, serviceName)
		if err := osutil.CreateDirs(path); err != nil {
			exit.Error(err)
		}
		//End
		// Complete the template execution transaction by copying the temporary dir to the target directory.
		if err := osutil.CopyRecursively(tempPath, svPath); err != nil {
			exit.Error(err)
		}
		exit.OK("Initialization complete")
	},
}

func getTemplatePath() (string, error) {
	tempPath := ""
	//Get GetExecutable Path
	execPath, _ := osutil.GetExecutablePath()
	gobxPath := filepath.Dir(execPath) + "/.."
	//Check if folder exists
	if exists, err := osutil.DirExists(gobxPath); !exists {
		return "", err
	}
	tempPath = gobxPath + "/example/"
	return tempPath, nil
}
