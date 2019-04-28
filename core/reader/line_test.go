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
	s.Equal(1, r.Row())
	s.Equal(1, r.Column())
	s.Equal(0, r.Absolute())
	r = NewPositionReader("thing1", Position{row: 4, column: 2})
	s.Equal(4, r.Row())
	s.Equal(2, r.Column())
	s.Equal(0, r.Absolute())
	r = NewPositionReader("thing1", Position{row: 4, column: 2, absolute: 42})
	s.Equal(4, r.Row())
	s.Equal(2, r.Column())
	s.Equal(42, r.Absolute())
	// Now create a reader with a position history
	r = NewPositionReader("thing1",
		Position{row: 1, column: 1, absolute: 0},
		Position{row: 1, column: 2, absolute: 1},
		Position{row: 1, column: 3, absolute: 2})
	s.Equal(1, r.Row())
	s.Equal(3, r.Column())
	s.Equal(2, r.Absolute())
}

func (s *LineReaderSuite) TestLastPositionIndex() {
	r := NewPositionReader("thing1",
		Position{row: 1, column: 1, absolute: 0},
		Position{row: 1, column: 2, absolute: 1},
		Position{row: 1, column: 3, absolute: 2})
	s.Equal(2, r.lastPositionIndex())
}

func (s *LineReaderSuite) TestLastPosition() {
	p1 := Position{row: 1, column: 1, absolute: 0}
	p2 := Position{row: 1, column: 2, absolute: 1}
	r := NewPositionReader("thing1", p1, p2)
	s.Equal(p2, r.lastPosition())
}

func (s *LineReaderSuite) TesDeletetLastPosition() {
	p1 := Position{row: 1, column: 1, absolute: 0}
	p2 := Position{row: 1, column: 2, absolute: 1}
	r := NewPositionReader("thing1", p1, p2)
	r.deleteLastPosition()
	s.Equal(p1, r.lastPosition())
}

func (s *LineReaderSuite) TestPushPosition() {
	p2 := Position{row: 1, column: 2, absolute: 1}
	r := NewPositionReader("thing1")
	r.pushPosition(p2)
	s.Equal(p2, r.lastPosition())
	s.Equal(2, len(r.positionStack))
}

func (s *LineReaderSuite) TestPushPositions() {
	p2 := Position{row: 1, column: 2, absolute: 1}
	p3 := Position{row: 1, column: 3, absolute: 2}
	p4 := Position{row: 2, column: 1, absolute: 3}
	r := NewPositionReader("thing1\nthing2")
	r.pushPositions(p2, p3, p4)
	s.Equal(p4, r.lastPosition())
	s.Equal(4, len(r.positionStack))
}

func (s *LineReaderSuite) TestPopPosition() {
	p1 := Position{row: 1, column: 1, absolute: 0}
	r := NewPositionReader("thing1")
	s.Equal(p1, r.popPosition())
	s.Equal(0, len(r.positionStack))
}

func (s *LineReaderSuite) TestNextRunePosition() {
	p2 := Position{row: 1, column: 1, absolute: 1}
	r := NewPositionReader("thing1")
	s.Equal(1, len(r.positionStack))
	s.Equal(p2, r.nextRunePosition())
	s.Equal(1, len(r.positionStack))
}

func (s *LineReaderSuite) TestReadRune() {
	r := NewPositionReader("thing1")
	s.Equal(1, r.Row())
	s.Equal(1, r.Column())
	s.Equal(0, r.Absolute())
	rn, sz, _ := r.ReadRune()
	s.Equal('t', rn)
	s.Equal(1, sz)
	s.Equal(1, r.Absolute())
	rn, sz, _ = r.ReadRune()
	s.Equal('h', rn)
	s.Equal(1, sz)
	s.Equal(2, r.Absolute())
}

func (s *LineReaderSuite) TestUneadRune() {
	r := NewPositionReader("thing1")
	_, _, _ = r.ReadRune()
	s.Equal(1, r.Absolute())
	_ = r.UnreadRune()
	s.Equal(0, r.Absolute())
	_, _, _ = r.ReadRune()
	_, _, _ = r.ReadRune()
	_, _, _ = r.ReadRune()
	s.Equal(3, r.Absolute())
	_ = r.UnreadRune()
	s.Equal(2, r.Absolute())
	_ = r.UnreadRune()
	s.Equal(1, r.Absolute())
	_ = r.UnreadRune()
	s.Equal(0, r.Absolute())
}
