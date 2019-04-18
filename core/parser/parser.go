package parser

import (
	"errors"

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

// ParseFromString parses code and returns a collection of nodes.
func ParseFromString(name, program string) []Node {
	return Parse(lexer.Lex(name, program))
}

// Parse takes lexed data and returns a tree of parsed nodes. This function is
//       a convenience wrapper around ParseAtoms.
func Parse(l *lexer.Lexer) []Node {
	return ParseAtoms(l, make([]Node, 0), ' ')
}

// ParseAtoms iterates over the items in the given lexed data and adds parsed
//            nodes to the given tree.
func ParseAtoms(l *lexer.Lexer, tree []Node, lookingFor rune) []Node {
	for item := l.NextAtom(); item.Type != lexer.AtomEOF; {
		node, err := ParseAtom(l, item, lookingFor)
		if err != nil {
			return tree
		}
		tree = append(tree, node)
		item = l.NextAtom()
	}
	return tree
}

// ParseAtom takes a given lexed item, determines its type, and then create a
//           corresponding and appropriate node for that atom.
func ParseAtom(l *lexer.Lexer, item lexer.Atom, lookingFor rune) (Node, error) {
	var node Node
	log.Tracef("Parsed: %#v", item)
	switch t := item.Type; t {
	case lexer.AtomIdent:
		node = NewIdentNode(item.Value)
	case lexer.AtomString:
		node = newStringNode(item.Value)
	case lexer.AtomInt:
		node = newIntNode(item.Value)
	case lexer.AtomFloat:
		node = newFloatNode(item.Value)
	case lexer.AtomComplex:
		node = newComplexNode(item.Value)
	case lexer.AtomLeftParen:
		node = newCallNode(ParseAtoms(l, make([]Node, 0), ')'))
	case lexer.AtomLeftVect:
		node = newVectNode(ParseAtoms(l, make([]Node, 0), ']'))
	case lexer.AtomRightParen:
		if lookingFor != ')' {
			log.Error(RightCurvedBracketError, item.Pos)
		}
		return nil, errors.New("done")
	case lexer.AtomRightVect:
		if lookingFor != ']' {
			log.Error(RightSquareBracketError, item.Pos)
		}
		return nil, errors.New("done")
	case lexer.AtomError:
		log.Panicf(UnspecifiedAtomError, lexer.AtomError, item.Pos)
	default:
		log.Panic(AtomTypeError)
	}
	return node, nil
}
