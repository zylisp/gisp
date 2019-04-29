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
	var l *LispReader
	l = NewLispReader("test-prog", "")
	l.Read()
	atoms := l.Atoms()
	s.Equal(1, len(atoms))
	s.Equal("AtomEOF", AtomName(atoms[0].Type))
}

func (s *LispReaderSuite) TestReadJustParens() {
	var l *LispReader
	l = NewLispReader("test-prog", "()")
	l.Read()
	atoms := l.Atoms()
	s.Equal(3, len(atoms))
	s.Equal("AtomLeftParen", AtomName(atoms[0].Type))
	s.Equal("AtomRightParen", AtomName(atoms[1].Type))
	s.Equal("AtomEOF", AtomName(atoms[2].Type))
}

func (s *LispReaderSuite) TestReadJustBrackets() {
	var l *LispReader
	l = NewLispReader("test-prog", "[]")
	l.Read()
	atoms := l.Atoms()
	s.Equal(3, len(atoms))
	s.Equal("AtomLeftVect", AtomName(atoms[0].Type))
	s.Equal("AtomRightVect", AtomName(atoms[1].Type))
	s.Equal("AtomEOF", AtomName(atoms[2].Type))
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
