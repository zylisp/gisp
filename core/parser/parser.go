package parser

import (
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/zylisp/zylisp/core/reader"
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
	r := reader.NewLispReader(name, program)
	r.Read()
	return Parse(r)
}

// Parse takes lexed data and returns a tree of parsed nodes. This function is
//       a convenience wrapper around ParseAtoms.
func Parse(l *reader.LispReader) []Node {
	return ParseAtoms(l, make([]Node, 0), ' ')
}

// ParseAtoms iterates over the items in the given lexed data and adds parsed
//            nodes to the given tree.
func ParseAtoms(l *reader.LispReader, tree []Node, lookingFor rune) []Node {
	for _, atom := range l.Atoms() {
		if atom.Type == reader.AtomEOF {
			break
		}
		node, err := ParseAtom(l, atom, lookingFor)
		if err != nil {
			return tree
		}
		tree = append(tree, node)
	}
	return tree
}

// ParseAtom takes a given Lisp atom, determines its type, and then create a
//           corresponding and appropriate node for that atom.
func ParseAtom(l *reader.LispReader, atom reader.Atom, lookingFor rune) (Node, error) {
	var node Node
	log.Tracef("Parsed: %#v", atom)
	switch t := atom.Type; t {
	case reader.AtomIdent:
		node = NewIdentNode(atom.Value)
	case reader.AtomString:
		node = newStringNode(atom.Value)
	case reader.AtomInt:
		node = newIntNode(atom.Value)
	case reader.AtomFloat:
		node = newFloatNode(atom.Value)
	case reader.AtomComplex:
		node = newComplexNode(atom.Value)
	case reader.AtomLeftParen:
		node = newCallNode(ParseAtoms(l, make([]Node, 0), ')'))
	case reader.AtomLeftVect:
		node = newVectNode(ParseAtoms(l, make([]Node, 0), ']'))
	case reader.AtomRightParen:
		if lookingFor != ')' {
			log.Error(RightCurvedBracketError, atom.Row(),
				atom.Column())
		}
		return nil, errors.New("done")
	case reader.AtomRightVect:
		if lookingFor != ']' {
			log.Error(RightSquareBracketError, atom.Row(),
				atom.Column())
		}
		return nil, errors.New("done")
	case reader.AtomError:
		log.Panicf(UnspecifiedAtomError, reader.AtomName(reader.AtomError),
			atom.Row(), atom.Column())
	default:
		log.Panic(AtomTypeError)
	}
	return node, nil
}
