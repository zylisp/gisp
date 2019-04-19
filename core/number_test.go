package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NumberSuite struct {
	suite.Suite
}

func TestNumberSuite(t *testing.T) {
	suite.Run(t, new(NumberSuite))
}

func (s *NumberSuite) TestMOD() {
	// s.Equal(MOD(4.0, 2.0), 0) // two floats
	// s.Equal(MOD(4, 2), 0)     // two ints
	// s.Equal(MOD(9.0, 3), 0)   // float, int
	s.Equal(MOD(7, 3.5), 1) // int, float
}

func (s *NumberSuite) TestADD() {
	// adding ints and floats
	s.Equal(ADD(3, 4.1), 7.1)
	s.Equal(ADD(4.1, 3), 7.1)
	// adding ints
	s.Equal(ADD(3, 4), 7.0)
	// NOTE ADD always returns a float
	// adding floats
	s.Equal(ADD(3.0, 4.1), 7.1)
	// s.Equal(func() { ADD("foo", 3) }, Panics, AddArgTypeError)
}

func (s *NumberSuite) TestAddNumbers() {
	// adding ints and floats
	sumNum, err := AddNumbers(Number{Value: 3}, Number{Value: 4.1})
	s.Nil(err)
	s.Equal(sumNum, Number{Value: 7.1})

	sumNum, err = AddNumbers(Number{Value: 4.1}, Number{Value: 3})
	s.Nil(err)
	s.Equal(sumNum, Number{Value: 7.1})

	sumNum, err = AddNumbers(Number{Value: 3}, Number{Value: 4})
	s.Nil(err)
	s.Equal(sumNum, Number{Value: 7})

	sumNum, err = AddNumbers(Number{Value: 4.1}, Number{Value: 3.0})
	s.Nil(err)
	s.Equal(sumNum, Number{Value: 7.1})
}

func (s *NumberSuite) TestIsInt() {
	s.True(IsInt(1))
	s.False(IsInt(4.0))
	s.True(IsInt(int64(3)))
}

func (s *NumberSuite) TestIsFloat() {
	s.False(IsFloat(1))
	s.True(IsFloat(4.0))
	s.True(IsFloat(3.5))
	s.True(IsFloat(float32(3.5)))
}

func (s *NumberSuite) TestToFloat() {
	x := Number{Value: 4}
	s.Equal(x.ToFloat(), 4.0)
	x.Value = 4.1
	s.Equal(x.ToFloat(), 4.1)
}

func (s *NumberSuite) TestToInt() {
	x := Number{Value: 4}
	s.Equal(x.ToInt(), 4)
	x.Value = 4.1
	s.Equal(x.ToInt(), 4)
}
