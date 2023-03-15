package a

import (
	"errors"
	"fmt"

	"go.uber.org/multierr"
)

func multierr_use() {
	err := multierr.Combine(
		nil, // successful request
		errors.New("call 2 failed"),
		errors.New("call 3 failed"),
	)
	err = multierr.Append(err, nil) // successful request
	err = multierr.Append(err, errors.New("call 5 failed"))

	errors := multierr.Errors(err) // want "CallExpr is here"
	for _, err := range errors {
		fmt.Println(err)
	}
}
