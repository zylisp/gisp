package reader

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AtomSuite struct {
	suite.Suite
}

func TestAtomSuite(t *testing.T) {
	suite.Run(t, new(AtomSuite))
}
func (s *AtomSuite) TestAtomName() {
	s.Equal(AtomName(0), "AtomError")
	s.Equal(AtomName(1), "AtomEOF")
	s.Equal(AtomName(2), "AtomLeftParen")
	s.Equal(AtomName(5), "AtomRightVect")
	s.Equal(AtomName(10), "AtomInt")
	s.Equal(AtomName(15), "AtomUnquoteSplice")
}
