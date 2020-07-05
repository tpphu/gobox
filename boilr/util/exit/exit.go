package exit

import (
	"fmt"
	"os"
	"github.com/tpphu/gobox/boilr/util/log"
)

const (
	// CodeOK indicates successful execution.
	CodeOK = 0

	// CodeError indicates erroneous execution.
	CodeError = 1

	// CodeFatal indicates erroneous use by user.
	CodeFatal = 2
)

// Fatal terminates execution using fatal exit code.
func Fatal(err error) {
	log.Fatal(fmt.Sprint(err))

	os.Exit(CodeFatal)
}

// Error terminates execution using unsuccessful execution exit code.
func Error(err error) {
	log.Error(err.Error())

	os.Exit(CodeError)
}

// GoodEnough terminates execution successfully, but indicates that something is missing.
func GoodEnough(fmtStr string, s ...interface{}) {
	log.Info(fmt.Sprintf(fmtStr, s...))
	os.Exit(CodeOK)
}

// OK terminates execution successfully.
func OK(fmtStr string, s ...interface{}) {
	log.Success(fmt.Sprintf(fmtStr, s...))
	os.Exit(CodeOK)
}
