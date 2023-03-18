package missingtypeguard

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "missingtypeguard checks if types that implement an interface have a type guard for it"

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

	typeGuardOwnersByInterfaces := typedMap[*typedMap[bool]]{}
	typesMap := typedMap[token.Pos]{}

	// find interfaces in the imported packages
	for _, pkg := range pass.Pkg.Imports() {
		for _, name := range pkg.Scope().Names() {
			obj, ok := pkg.Scope().Lookup(name).(*types.TypeName)
			if !ok {
				continue
			}

			if named, ok := obj.Type().(*types.Named); ok {
				if _, ok := named.Underlying().(*types.Interface); ok {
					typeGuardOwnersByInterfaces.Set(named, &typedMap[bool]{})
				}
			}
		}
	}

	// find interfaces and type guards in the current package
	inspect.Preorder([]ast.Node{
		(*ast.TypeSpec)(nil),
		(*ast.ValueSpec)(nil),
	}, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.TypeSpec:
			switch n.Type.(type) {
			case *ast.InterfaceType:
				if typeGuardOwnersByInterfaces.At(pass.TypesInfo.TypeOf(n.Name)) == nil {
					typeGuardOwnersByInterfaces.Set(pass.TypesInfo.TypeOf(n.Name), &typedMap[bool]{})
				}

			default:
				typesMap.Set(pass.TypesInfo.TypeOf(n.Name), n.Pos())
			}

		case *ast.ValueSpec:
			if n.Type == nil || len(n.Values) != 1 {
				return
			}

			itype := pass.TypesInfo.TypeOf(n.Type)
			if typeGuardOwnersByInterfaces.At(itype) == nil {
				typeGuardOwnersByInterfaces.Set(itype, &typedMap[bool]{})
			}

			typeGuardOwnersByInterfaces.At(itype).Set(pass.TypesInfo.TypeOf(n.Values[0]), true)
		}
	})

	// check if types that implement an interface have a type guard for it
	typesMap.Iterate(func(ntype types.Type, pos token.Pos) {
		typeGuardOwnersByInterfaces.Iterate(func(itype types.Type, typeGuardOwners *typedMap[bool]) {
			i, ok := itype.Underlying().(*types.Interface)
			if !ok {
				return
			}

			if types.Implements(ntype, i) {
				if !typeGuardOwners.At(ntype) {
					pass.Reportf(pos, "%s is missing a type guard for %s", ntype, itype)
				}

				return // no need to check the pointer
			}

			nptype := types.NewPointer(ntype)
			if types.Implements(nptype, i) && !typeGuardOwners.At(nptype) {
				pass.Reportf(pos, "the pointer of %s is missing a type guard for %s", ntype, itype)
			}

		})
	})

	return nil, nil
}
