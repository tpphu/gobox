package cmd

import (
	"fmt"
	"github.com/tpphu/gobox/boilr/cmd/util"
	boilr "github.com/tpphu/gobox/boilr/configuration"
	"github.com/tpphu/gobox/boilr/util/exit"
	"github.com/tpphu/gobox/boilr/util/validate"
)

// MustValidateVarArgs validates given variadic arguments with the supplied validation function.
// If there are any errors it exits the execution.
func MustValidateVarArgs(args []string, v validate.Argument) {
	if err := util.ValidateVarArgs(args, v); err != nil {
		exit.Error(err)
	}
}

// MustValidateArgs validates given arguments with the supplied validation functions.
// If there are any errors it exits the execution.
func MustValidateArgs(args []string, validations []validate.Argument) {
	if err := util.ValidateArgs(args, validations); err != nil {
		exit.Error(err)
	}
}

// MustValidateTemplateDir ensures that template directory is initialized.
func MustValidateTemplateDir() {
	isInitialized, err := boilr.IsTemplateDirInitialized()
	if err != nil {
		exit.Error(err)
	}

	if !isInitialized {
		exit.Error(fmt.Errorf("Template registry is not initialized. Please run `init` command to initialize it."))
	}
}
