package a

import "go.uber.org/multierr"

var _ = multierr.Errors

type Spurious_Multierr struct {
	Errors func(error) error
}

func multierr_spurious() {
	multierr := Spurious_Multierr{}
	multierr.Errors(nil) // ok
}
