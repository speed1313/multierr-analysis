package a

import (
	"errors"
	"fmt"

	me "go.uber.org/multierr"
)

func multierr_alias() {
	err := me.Combine(
		nil, // successful request
		errors.New("call 2 failed"),
		errors.New("call 3 failed"),
	)
	err = me.Append(err, nil) // successful request
	err = me.Append(err, errors.New("call 5 failed"))

	errors := me.Errors(err) // want "CallExpr is here"
	for _, err := range errors {
		fmt.Println(err)
	}
}
