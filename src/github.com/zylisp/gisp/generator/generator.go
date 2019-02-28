package generator

import (
	"github.com/zylisp/gisp/parser"
	"fmt"
	"go/ast"
	"go/token"
)

var anyType = makeSelectorExpr(ast.NewIdent("core"), ast.NewIdent("Any"))

func GenerateAST(tree []parser.Node) *ast.File {
	f := &ast.File{Name: ast.NewIdent("main")}
	decls := make([]ast.Decl, 0, len(tree))

	if len(tree) < 1 {
		return f
	}

	// you can only have (ns ...) as the first form
	if isNSDecl(tree[0]) {
		name, imports := getNamespace(tree[0].(*parser.CallNode))

		f.Name = name
		if imports != nil {
			decls = append(decls, imports)
		}

		tree = tree[1:]
	}

	decls = append(decls, generateDecls(tree)...)

	f.Decls = decls
	return f
}

func generateDecls(tree []parser.Node) []ast.Decl {
	decls := make([]ast.Decl, len(tree))

	for i, node := range tree {
		if node.Type() != parser.NodeCall {
			panic("expected call node in root scope!")
		}

		decls[i] = evalDeclNode(node.(*parser.CallNode))
	}

	return decls
}

func evalDeclNode(node *parser.CallNode) ast.Decl {
	// Let's just assume that all top-level functions called will be "def"
	if node.Callee.Type() != parser.NodeIdent {
		panic("expecting call to identifier (i.e. def, defconst, etc.)")
	}

	callee := node.Callee.(*parser.IdentNode)
	switch callee.Ident {
	case "def":
		return evalDef(node)
	}

	return nil
}

func evalDef(node *parser.CallNode) ast.Decl {
	if len(node.Args) < 2 {
		panic(fmt.Sprintf("expecting expression to be assigned to variable: %q", node.Args[0]))
	}

	val := EvalExpr(node.Args[1])
	fn, ok := val.(*ast.FuncLit)

	ident := makeIdomaticIdent(node.Args[0].(*parser.IdentNode).Ident)

	if ok {
		if ident.Name == "main" {
			mainable(fn)
		}

		return makeFunDeclFromFuncLit(ident, fn)
	} else {
		return makeGeneralDecl(token.VAR, []ast.Spec{makeValueSpec([]*ast.Ident{ident}, []ast.Expr{val}, nil)})
	}
}

func isNSDecl(node parser.Node) bool {
	if node.Type() != parser.NodeCall {
		return false
	}

	call := node.(*parser.CallNode)
	if call.Callee.(*parser.IdentNode).Ident != "ns" {
		return false
	}

	if len(call.Args) < 1 {
		return false
	}

	return true
}

func getNamespace(node *parser.CallNode) (*ast.Ident, ast.Decl) {
	return getPackageName(node), getImports(node)
}

func getPackageName(node *parser.CallNode) *ast.Ident {
	if node.Args[0].Type() != parser.NodeIdent {
		panic("ns package name is not an identifier!")
	}

	return ast.NewIdent(node.Args[0].(*parser.IdentNode).Ident)
}

func checkNSArgs(node *parser.CallNode) bool {
	if node.Callee.Type() != parser.NodeIdent {
		return false
	}

	if callee := node.Callee.(*parser.IdentNode); callee.Ident != "ns" {
		return false
	}

	return true
}
