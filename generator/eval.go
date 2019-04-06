package generator

import (
	"go/ast"
	"go/token"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/common"
	"github.com/zylisp/zylisp/parser"
)

// EvalExprs is a function that takes a collection of parser nodes and returns
// a collection of AST expressions.
func EvalExprs(nodes []parser.Node) []ast.Expr {
	out := make([]ast.Expr, len(nodes))

	for i, node := range nodes {
		out[i] = EvalExpr(node)
	}

	return out
}

// EvalExpr is a function that takes a parser node and returns an AST
// expression.
func EvalExpr(node parser.Node) ast.Expr {
	t := node.Type()
	log.Trace("Evaluating node:", node, "of type:", t)
	log.Trace("Node data:", node)
	switch t {
	case parser.NodeCall:
		node := node.(*parser.CallNode)
		return evalFuncCall(node)

	case parser.NodeVector:
		node := node.(*parser.VectorNode)
		return makeVector(anyType, EvalExprs(node.Nodes))

	case parser.NodeNumber:
		node := node.(*parser.NumberNode)
		return makeBasicLit(node.NumberType, node.Value)

	case parser.NodeString:
		node := node.(*parser.StringNode)
		return makeBasicLit(token.STRING, node.Value)

	case parser.NodeIdent:
		node := node.(*parser.IdentNode)
		return makeIdomaticSelector(node.Ident)

	default:
		log.Panicf("%s: %#v", common.NotImplementedError, t)
		return nil
	}
}
