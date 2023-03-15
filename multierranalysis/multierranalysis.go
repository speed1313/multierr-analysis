package multierranalysis

import (
	"go/ast"

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

	nodeFilter := []ast.Node{
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ImportSpec:
			if n.Path.Value == `"go.uber.org/multierr"` {
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

	return nil, nil
}
