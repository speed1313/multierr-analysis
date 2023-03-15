# multierr-analysis
multierr static analysis 

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
