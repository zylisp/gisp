package generator

import (
	"go/ast"
	"strconv"
	"strings"

	"github.com/zylisp/zylisp/common"
	// "github.com/zylisp/zylisp/core/parser"
)

// XXX Currently unused; remove?
// func makeIdentSlice(nodes []*parser.IdentNode) []*ast.Ident {
// 	out := make([]*ast.Ident, len(nodes))
// 	for i, node := range nodes {
// 		out[i] = ast.NewIdent(node.Ident)
// 	}
// 	return out
// }

func makeSelectorExpr(x ast.Expr, sel *ast.Ident) *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X:   x,
		Sel: sel,
	}
}

func makeIdomaticSelector(src string) ast.Expr {
	strs := strings.Split(src, "/")
	var expr ast.Expr = makeIdomaticIdent(strs[0])

	for i := 1; i < len(strs); i++ {
		ido := common.CamelCase(strs[i], true)
		expr = makeSelectorExpr(expr, ast.NewIdent(ido))
	}

	return expr
}

func makeIdomaticIdent(src string) *ast.Ident {
	if src == "_" {
		return ast.NewIdent(src)
	}
	return ast.NewIdent(common.CamelCase(src, false))
}

var gensyms = func() <-chan string {
	syms := make(chan string)
	go func() {
		i := 0
		for {
			syms <- "GEN" + strconv.Itoa(i)
			i++
		}
	}()
	return syms
}()

func generateIdent() *ast.Ident {
	return ast.NewIdent(<-gensyms)
}
