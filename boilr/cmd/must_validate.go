package cmd

import (
	"github.com/tpphu/gobox/boilr/cmd/util"
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
