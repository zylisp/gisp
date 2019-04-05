package parser

import (
	"go/token"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/lexer"
)

type Pos int

func (p Pos) Position() Pos {
	return p
}

var nilNode = NewIdentNode("nil")

func ParseFromString(name, program string) []Node {
	return Parse(lexer.Lex(name, program))
}

func Parse(l *lexer.Lexer) []Node {
	return parser(l, make([]Node, 0), ' ')
}

func parser(l *lexer.Lexer, tree []Node, lookingFor rune) []Node {
	for item := l.NextAtom(); item.Type != lexer.AtomEOF; {
		log.Tracef("Parsed: %s", item)
		switch t := item.Type; t {
		case lexer.AtomIdent:
			tree = append(tree, NewIdentNode(item.Value))
		case lexer.AtomString:
			tree = append(tree, newStringNode(item.Value))
		case lexer.AtomInt:
			tree = append(tree, newIntNode(item.Value))
		case lexer.AtomFloat:
			tree = append(tree, newFloatNode(item.Value))
		case lexer.AtomComplex:
			tree = append(tree, newComplexNode(item.Value))
		case lexer.AtomLeftParen:
			tree = append(tree, newCallNode(parser(l, make([]Node, 0), ')')))
		case lexer.AtomLeftVect:
			tree = append(tree, newVectNode(parser(l, make([]Node, 0), ']')))
		case lexer.AtomRightParen:
			if lookingFor != ')' {
				log.Error(RightCurvedBracketError, item.Pos)
			}
			return tree
		case lexer.AtomRightVect:
			if lookingFor != ']' {
				log.Error(RightSquareBracketError, item.Pos)
			}
			return tree
		case lexer.AtomError:
			log.Panicf(UnspecifiedAtomError, lexer.AtomError, item.Pos)
		default:
			log.Panic(AtomTypeError)
		}
		item = l.NextAtom()
	}
	return tree
}

func NewIdentNode(name string) *IdentNode {
	return &IdentNode{NodeType: NodeIdent, Ident: name}
}

func newStringNode(val string) *StringNode {
	return &StringNode{NodeType: NodeString, Value: val}
}

func newIntNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.INT}
}

func newFloatNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.FLOAT}
}

func newComplexNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.IMAG}
}

// We return Node here, because it could be that it's nil
func newCallNode(args []Node) Node {
	if len(args) > 0 {
		return &CallNode{NodeType: NodeCall, Callee: args[0], Args: args[1:]}
	} else {
		return nilNode
	}
}

func newVectNode(content []Node) *VectorNode {
	return &VectorNode{NodeType: NodeVector, Nodes: content}
}
