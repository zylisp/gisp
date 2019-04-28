package line

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LineReaderSuite struct {
	suite.Suite
}

func TestLineReaderSuite(t *testing.T) {
	suite.Run(t, new(LineReaderSuite))
}

func (s *LineReaderSuite) TestNewPositionReader() {
	r := NewPositionReader("thing1")
	s.Equal(1, r.row)
	s.Equal(1, r.column)
	s.Equal(0, r.absolute)
	r = NewPositionReader("thing1", PositionOpts{row: 4, column: 2})
	s.Equal(4, r.row)
	s.Equal(2, r.column)
	s.Equal(0, r.absolute)
	r = NewPositionReader("thing1", PositionOpts{row: 4, column: 2, absolute: 42})
	s.Equal(4, r.row)
	s.Equal(2, r.column)
	s.Equal(42, r.absolute)
}

func (s *LineReaderSuite) TestReadRune() {
	r := NewPositionReader("thing1")
	s.Equal(1, r.row)
	s.Equal(1, r.column)
	s.Equal(0, r.absolute)
	rn, sz, _ := r.ReadRune()
	s.Equal('t', rn)
	s.Equal(1, sz)
	s.Equal(1, r.absolute)
}

func (s *LineReaderSuite) TestUneadRune() {
	r := NewPositionReader("thing1")
	_, _, _ = r.ReadRune()
	s.Equal(1, r.absolute)
	_ = r.UnreadRune()
	s.Equal(0, r.absolute)
	_, _, _ = r.ReadRune()
	_, _, _ = r.ReadRune()
	_, _, _ = r.ReadRune()
	s.Equal(3, r.absolute)
	_ = r.UnreadRune()
	s.Equal(2, r.absolute)
	_ = r.UnreadRune()
	s.Equal(1, r.absolute)
	_ = r.UnreadRune()
	s.Equal(0, r.absolute)
}
