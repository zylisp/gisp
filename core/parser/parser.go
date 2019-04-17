package parser

import (
	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/core/lexer"
)

// Pos type
type Pos int

// Position returns a pos type
func (p Pos) Position() Pos {
	return p
}

var nilNode = NewIdentNode("nil")

// ParseFromString parses code and returns a collection of nodes
func ParseFromString(name, program string) []Node {
	return Parse(lexer.Lex(name, program))
}

// Parse takes lexed data and returns a tree of parsed nodes.
func Parse(l *lexer.Lexer) []Node {
	return parser(l, make([]Node, 0), ' ')
}

func parser(l *lexer.Lexer, tree []Node, lookingFor rune) []Node {
	for item := l.NextAtom(); item.Type != lexer.AtomEOF; {
		log.Tracef("Parsed: %#v", item)
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
