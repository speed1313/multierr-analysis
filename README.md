# multierr-analysis

multierr-analysis is a static analysis to check if the multierr.Errors is used.

- Example

```
import "go.uber.org/multierr"

...

errors := multierr.Errors(err) // want "CallExpr is here"
```

