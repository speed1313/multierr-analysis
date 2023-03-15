package main

import (
	"github.com/speed1313/multierranalysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(multierranalysis.Analyzer) }
