package missingtypeguard

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "missingtypeguard is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "missingtypeguard",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	typeGuardOwners := make(map[types.Type]struct{})

	// 1. Find all type guards for `Animal`
	inspect.Preorder([]ast.Node{(*ast.ValueSpec)(nil)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ValueSpec:
			if n.Type == nil || len(n.Values) != 1 {
				return
			}

			if ident, ok := n.Type.(*ast.Ident); ok && ident.Name == "Animal" {
				switch concreteType := n.Values[0].(type) {
				// var _ Animal = dog{}
				case *ast.CompositeLit:
					typeGuardOwners[pass.TypesInfo.TypeOf(concreteType.Type)] = struct{}{}
				}
			}
		}
	})

	// 2. Find structs missing type guards for `Animal`
	inspect.Preorder([]ast.Node{(*ast.TypeSpec)(nil)}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			if _, ok := n.Type.(*ast.StructType); !ok {
				return
			}

			if _, ok := typeGuardOwners[pass.TypesInfo.TypeOf(n.Name)]; !ok {
				pass.Reportf(n.Pos(), "%s is missing a type guard for Animal", n.Name.Name)
			}
		}
	})

	return nil, nil
}
