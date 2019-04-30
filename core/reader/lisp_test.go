package reader

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LispReaderSuite struct {
	suite.Suite
}

func TestLispReaderSuite(t *testing.T) {
	suite.Run(t, new(LispReaderSuite))
}

func (s *LispReaderSuite) TestReadEmpty() {
	l := NewLispReader("test-prog", "")
	l.Read()
	atoms := l.Atoms()
	s.Equal(1, len(atoms))
	s.Equal("AtomEOF", AtomName(atoms[0].Type))
}

func (s *PositionReaderSuite) TestRead1EOF() {
	var rn rune
	l := NewLispReader("test-prog", "()")
	// This will bring us to the end of the string:
	rn = l.read1()
	s.Equal("(", string(rn))
	rn = l.read1()
	s.Equal(")", string(rn))
	// This should return our rune-safe EOF value:
	rn = l.read1()
	s.Equal(EOF, rn)
}

func (s *LispReaderSuite) TestReadJustParens() {
	l := NewLispReader("test-prog", "()")
	l.Read()
	atoms := l.Atoms()
	s.Equal(3, len(atoms))
	s.Equal("AtomLeftParen", AtomName(atoms[0].Type))
	s.Equal("AtomRightParen", AtomName(atoms[1].Type))
	s.Equal("AtomEOF", AtomName(atoms[2].Type))
}

func (s *LispReaderSuite) TestReadJustBrackets() {
	l := NewLispReader("test-prog", "[]")
	l.Read()
	atoms := l.Atoms()
	s.Equal(3, len(atoms))
	s.Equal("AtomLeftVect", AtomName(atoms[0].Type))
	s.Equal("AtomRightVect", AtomName(atoms[1].Type))
	s.Equal("AtomEOF", AtomName(atoms[2].Type))
}

func (s *LispReaderSuite) TestReadJustString() {
	l := NewLispReader("test-prog", `"space dog"`)
	l.Read()
	atoms := l.Atoms()
	s.Equal(2, len(atoms))
	s.Equal("AtomString", AtomName(atoms[0].Type))
	s.Equal("space dog", atoms[0].Value)
	s.Equal("AtomEOF", AtomName(atoms[1].Type))
}

func (s *LispReaderSuite) TestReadEmptyString() {
	l := NewLispReader("test-prog", `""`)
	l.Read()
	atoms := l.Atoms()
	s.Equal(2, len(atoms))
	s.Equal("AtomString", AtomName(atoms[0].Type))
	s.Equal("", atoms[0].Value)
	s.Equal("AtomEOF", AtomName(atoms[1].Type))
}

func (s *LispReaderSuite) TestReadStringWithEscapes() {
	l := NewLispReader("test-prog", "\"space dog\"")
	l.Read()
	atoms := l.Atoms()
	s.Equal(2, len(atoms))
	s.Equal("AtomString", AtomName(atoms[0].Type))
	s.Equal("space dog", atoms[0].Value)
	s.Equal("AtomEOF", AtomName(atoms[1].Type))
	// XXX Add more tests:
	// "space\"dog" -> s.Equal(`space"dog`, atoms[0].Value)
	// "space\\dog" -> s.Equal("space\dog", atoms[0].Value)
	// "\"space dog\"" -> s.Equal(`"space dog"``, atoms[0].Value)
	// "space\dog" -> s.Equal("spacedog", atoms[0].Value)
}

func (s *LispReaderSuite) TestReadComments() {
	var l *LispReader
	var atoms []Atom
	l = NewLispReader("test-prog", ";")
	l.Read()
	atoms = l.Atoms()
	s.Equal(1, len(atoms))
	s.Equal(";", atoms[0].Value)
	s.Equal("AtomEOF", AtomName(atoms[0].Type))
	l = NewLispReader("test-prog", `"space dog" ; comment`)
	l.Read()
	atoms = l.Atoms()
	s.Equal(2, len(atoms))
	s.Equal("AtomString", AtomName(atoms[0].Type))
	s.Equal("AtomEOF", AtomName(atoms[1].Type))
	l = NewLispReader("test-prog", "(); comment 1\n; comment 2\n()")
	l.Read()
	atoms = l.Atoms()
	s.Equal(5, len(atoms))
	s.Equal("AtomLeftParen", AtomName(atoms[0].Type))
	s.Equal("AtomRightParen", AtomName(atoms[1].Type))
	s.Equal("AtomLeftParen", AtomName(atoms[2].Type))
	s.Equal("AtomRightParen", AtomName(atoms[3].Type))
	s.Equal("AtomEOF", AtomName(atoms[4].Type))
}
