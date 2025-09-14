// Package errors:
package errors

import "fmt"

type EnvError struct {
	property string
}

func (err *EnvError) Error() string {
	return fmt.Sprintf("env variable: %s is undefined or invalid", err.property)
}

func NewEnvError(property string) error {
	return &EnvError{property}
}
