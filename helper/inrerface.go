package helper

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type interface_config struct {
	filePath string
}

func NewInterfaceConfig(filePath string) *interface_config {
	return &interface_config{
		filePath: filePath,
	}
}

// function to check if the file contains an interface
func (c *interface_config) Contains() bool {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, c.filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return false
	}

	for _, decl := range node.Decls {
		if genDecl, isGenDecl := decl.(*ast.GenDecl); isGenDecl {
			if genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, isTypeSpec := spec.(*ast.TypeSpec); isTypeSpec {
						if _, isInterface := typeSpec.Type.(*ast.InterfaceType); isInterface {
							return true
						}
					}
				}
			}
		}
	}

	return false
}
