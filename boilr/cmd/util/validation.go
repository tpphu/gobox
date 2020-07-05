package util

import (
	"errors"
	"fmt"
	"github.com/tpphu/gobox/boilr/util/validate"
)

var (
	// ErrUnexpectedArgs indicates that the given number of arguments exceed the expected number of arguments.
	ErrUnexpectedArgs = errors.New("unexpected arguments")

	// ErrNotEnoughArgs indicates that the given number of arguments does not match the expected number of arguments.
	ErrNotEnoughArgs = errors.New("not enough arguments")
)

const (
	// InvalidArg error message format string for filling in the details of an invalid arg.
	InvalidArg = "invalid argument for %q: %q, should be a valid %v"
)

// ValidateArgCount validates the number of arguments and returns the corresponding error if there are any.
func ValidateArgCount(expectedArgNo, argNo int) error {
	switch {
	case expectedArgNo < argNo:
		return ErrUnexpectedArgs
	case expectedArgNo > argNo:
		return ErrNotEnoughArgs
	case expectedArgNo == argNo:
	}

	return nil
}

// ValidateVarArgs validates variadic arguments with the given validate.Argument function.
func ValidateVarArgs(args []string, v validate.Argument) error {
	if len(args) == 0 {
		return ErrNotEnoughArgs
	}

	for _, arg := range args {
		if ok := v.Validate(arg); !ok {
			return fmt.Errorf(InvalidArg, v.Name, arg, v.Validate.TypeName())
		}
	}

	return nil
}

// ValidateArgs validates arguments with the given validate.Argument functions.
// Two arguments must contain the same number of elements.
func ValidateArgs(args []string, validations []validate.Argument) error {
	if err := ValidateArgCount(len(validations), len(args)); err != nil {
		return err
	}

	for i, arg := range validations {
		if ok := arg.Validate(args[i]); !ok {
			return fmt.Errorf(InvalidArg, arg.Name, args[i], arg.Validate.TypeName())
		}
	}

	return nil
}
