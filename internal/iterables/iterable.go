package iterables

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/ufukty/gonfique/internal/files"
	"github.com/ufukty/gonfique/internal/namings"
)

func DetectIterators(f *files.File) error {
	fds := []*ast.FuncDecl{}
	gds := []*ast.GenDecl{
		f.Isolated,
		{ // temporary
			Tok:   token.TYPE,
			Specs: []ast.Spec{&ast.TypeSpec{Name: ast.NewIdent("Config"), Type: f.Cfg}}},
	}
	for _, gd := range gds {
		for _, s := range gd.Specs {
			if ts, ok := s.(*ast.TypeSpec); ok {
				if st, ok := ts.Type.(*ast.StructType); ok {
					var cti *ast.Ident
					for _, f := range st.Fields.List {
						if ti, ok := f.Type.(*ast.Ident); ok {
							if cti == nil {
								cti = ti
								continue
							} else if ti.Name == cti.Name {
								continue
							}
						}
						cti = nil
						break
					}
					// if the all fields have same Ident in their types;
					// generate a FuncDecl which its body consists by a ReturnStmt of map[string]cti
					// the map has the exact same amount of Fields struct type has
					if cti != nil {
						elements := []ast.Expr{}
						for _, f := range st.Fields.List {
							keyname, err := namings.StripKeyname(f.Tag.Value)
							if err != nil {
								return fmt.Errorf("could not strip the keyname in %s.%s field tag list: %w", ts.Name.Name, f.Names[0].Name, err)
							}
							elements = append(elements, &ast.KeyValueExpr{
								Key:   &ast.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", keyname)},
								Value: &ast.SelectorExpr{X: &ast.Ident{Name: "a"}, Sel: f.Names[0]},
							})
						}
						fds = append(fds, &ast.FuncDecl{
							Recv: &ast.FieldList{List: []*ast.Field{{Names: []*ast.Ident{{Name: "a"}}, Type: ts.Name}}},
							Name: &ast.Ident{Name: "Range"},
							Type: &ast.FuncType{
								Params:  &ast.FieldList{},
								Results: &ast.FieldList{List: []*ast.Field{{Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: cti}}}},
							},
							Body: &ast.BlockStmt{List: []ast.Stmt{&ast.ReturnStmt{
								Results: []ast.Expr{&ast.CompositeLit{
									Type: &ast.MapType{Key: &ast.Ident{Name: "string"}, Value: cti},
									Elts: elements,
								}},
							}}},
						})
					}
				}
			}
		}
	}
	f.Iterators = fds
	return nil
}
