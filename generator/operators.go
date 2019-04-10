package generator

import (
	"go/ast"
	"go/token"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/core/parser"
)

var (
	callableOperators = []string{">", ">=", "<", "<=", "=", "+", "-", "*", "/", "mod"}
	logicOperatorMap  = map[string]token.Token{
		"and": token.LAND,
		"or":  token.LOR,
	}

	unaryOperatorMap = map[string]token.Token{
		"!": token.NOT,
	}
)

func isCallableOperator(node *parser.CallNode) bool {
	if node.Callee.Type() != parser.NodeIdent {
		return false
	}

	ident := node.Callee.(*parser.IdentNode).Ident

	return isInSlice(ident, callableOperators)
}

// We handle comparisons as a call to some go code, since you can only
// compare ints, floats, cmplx, and such, you know...
// We handle arithmetic operations as function calls, since all args are evaluated
func makeNAryCallableExpr(node *parser.CallNode) *ast.CallExpr {
	op := node.Callee.(*parser.IdentNode).Ident
	args := EvalExprs(node.Args)
	var selector string

	// TODO: abstract this away into a map!!!
	switch op {
	case ">":
		selector = "GT"
	case ">=":
		selector = "GTEQ"
	case "<":
		selector = "LT"
	case "<=":
		selector = "LTEQ"
	case "=":
		selector = "EQ"
	case "+":
		selector = "ADD"
	case "-":
		selector = "SUB"
	case "*":
		selector = "MUL"
	case "/":
		selector = "DIV"
	case "mod":
		argsCount := len(node.Args)
		requiredCount := 2
		if argsCount > requiredCount {
			log.Error(TooManyArgsError, "mod", requiredCount)
		}
		selector = "MOD"
	}
	return makeFuncCall(makeSelectorExpr(ast.NewIdent("core"), ast.NewIdent(selector)), args)
}

func isLogicOperator(node *parser.CallNode) bool {
	if node.Callee.Type() != parser.NodeIdent {
		return false
	}

	_, ok := logicOperatorMap[node.Callee.(*parser.IdentNode).Ident]
	argsCount := len(node.Args)
	requiredCount := 2
	if argsCount < requiredCount && ok {
		log.Error(TooManyArgsError, "a logical operator", requiredCount)
	}
	return ok
}

// But logical comparisons are done properly, since those can short-circuit
func makeNAryLogicExpr(node *parser.CallNode) *ast.BinaryExpr {
	op := logicOperatorMap[node.Callee.(*parser.IdentNode).Ident]
	outer := makeBinaryExpr(op, EvalExpr(node.Args[0]), EvalExpr(node.Args[1]))

	for i := 2; i < len(node.Args); i++ {
		outer = makeBinaryExpr(op, outer, EvalExpr(node.Args[i]))
	}
	return outer
}

func makeBinaryExpr(op token.Token, x, y ast.Expr) *ast.BinaryExpr {
	return &ast.BinaryExpr{
		X:  x,
		Y:  y,
		Op: op,
	}
}

func isUnaryOperator(node *parser.CallNode) bool {
	if node.Callee.Type() != parser.NodeIdent {
		return false
	}

	_, ok := unaryOperatorMap[node.Callee.(*parser.IdentNode).Ident]

	if len(node.Args) != 1 && ok {
		log.Error(UnaryArgsCountError)
	}

	return ok
}

func makeUnaryExpr(op token.Token, x ast.Expr) *ast.UnaryExpr {
	return &ast.UnaryExpr{
		X:  x,
		Op: op,
	}
}

func isInSlice(elem string, slice []string) bool {
	for _, el := range slice {
		if elem == el {
			return true
		}
	}

	return false
}
