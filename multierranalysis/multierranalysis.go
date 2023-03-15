package multierranalysis

import (
	"go/ast"
	"strconv"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "multierranalysis is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "multierranalysis",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	var rerr error
	//var mp map[string]string

	inspect.Preorder(nil, func(n ast.Node) {
		if rerr != nil {
			return
		}
		switch n := n.(type) {
		case *ast.ImportSpec:
			value, err := strconv.Unquote(n.Path.Value)
			if err != nil {
				rerr = err
				return
			}

			if value == "go.uber.org/multierr" {
				if n.Name != nil {
					//importName = n.Name.Name
				}
				pass.Reportf(n.Pos(), "multierr is imported")
			}
		case *ast.Ident:
			if n.Name == "Errors" {
				pass.Reportf(n.Pos(), "Errors is here")
			}
		case *ast.CallExpr:
			// TODO: check
			pass.Reportf(n.Pos(), "CallExpr is here")
		}
	})

	return nil, rerr
}
