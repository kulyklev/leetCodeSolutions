// Package solutions contains all solved algorithm tasks
package solutions

import "errors"

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")
var ErrFailedToRun = errors.New("failed to run command")

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
	Description() string
}
