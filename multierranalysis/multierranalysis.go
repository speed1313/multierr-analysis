package multierranalysis

import (
	"go/ast"

	"github.com/gostaticanalysis/analysisutil"
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
	pkgs := pass.Pkg.Imports()
	obj := analysisutil.LookupFromImports(pkgs, "go.uber.org/multierr", "Errors")
	if obj == nil {
		return nil, nil
	}
	types := pass.TypesInfo

	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			value, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}
			if obj == types.ObjectOf(value.Sel) {
				pass.Reportf(n.Pos(), "CallExpr is here")
			}
		}
	})

	return nil, nil
}
