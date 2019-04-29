package reader

import log "github.com/sirupsen/logrus"

// AtomType atom type
type AtomType int

// Atom object
type Atom struct {
	Type     AtomType
	Value    string
	Position Position
}

// NewAtom ...
func NewAtom(l *LispReader, atomType AtomType) Atom {
	pos := l.lastPosition()
	val := l.ReadToken()
	log.Tracef("Creating atom of type %s from token %s at position %#v ...",
		AtomName(atomType), val, pos)
	return Atom{Type: atomType, Value: val, Position: pos}
}

// Lexer atom constants
const (
	AtomError AtomType = iota
	AtomEOF

	AtomLeftParen
	AtomRightParen
	AtomLeftVect
	AtomRightVect

	AtomIdent
	AtomString
	AtomChar
	AtomFloat
	AtomInt
	AtomComplex

	// Ok, these aren't really atoms but ...
	AtomQuote
	AtomQuasiQuote
	AtomUnquote
	AtomUnquoteSplice
)

var atomNames = []string{
	AtomError:         "AtomError",
	AtomEOF:           "AtomEOF",
	AtomLeftParen:     "AtomLeftParen",
	AtomRightParen:    "AtomRightParen",
	AtomLeftVect:      "AtomLeftVect",
	AtomRightVect:     "AtomRightVect",
	AtomIdent:         "AtomIdent",
	AtomString:        "AtomString",
	AtomChar:          "AtomChar",
	AtomFloat:         "AtomFloat",
	AtomInt:           "AtomInt",
	AtomComplex:       "AtomComplex",
	AtomQuote:         "AtomQuote",
	AtomQuasiQuote:    "AtomQuasiQuote",
	AtomUnquote:       "AtomUnquote",
	AtomUnquoteSplice: "AtomUnquoteSplice",
}

// Column returns the column value of the associated item Position
func (a *Atom) Column() int {
	return a.Position.column
}

// Row returns the row value of the associated item Position
func (a *Atom) Row() int {
	return a.Position.row
}

// Absolute returns the absolute value of the associated item Position
func (a *Atom) Absolute() int {
	return a.Position.absolute
}

// AtomName returns the name of the atom for a given atom value
func AtomName(atomType AtomType) string {
	return atomNames[atomType]
}
