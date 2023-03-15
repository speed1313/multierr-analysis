# multierr-analysis
multierr-analysis is a static analysis to check if the multierr.Errors is used.

# install

You can get `multierr-analysis` by go install command. 

```
$ go install github.com/speed1313/multierr-analysis
```

# How to use

`multierr-analysis` run with go vet.

```
$ go vet -vettool=$(multierr-analysis) ./..
```

# Example

```
import "go.uber.org/multierr"

...

errors := multierr.Errors(err) // want "CallExpr is here"
```

