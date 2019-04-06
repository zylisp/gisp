package core

import (
	"testing"

	. "github.com/masukomi/check"
)

// hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }

type CoreSuite struct{}

var _ = Suite(&CoreSuite{})

func (s *CoreSuite) Test_IsInt(c *C) {
	c.Assert(IsInt(1), IsTrue)
	c.Assert(IsInt(4.0), IsFalse)
	c.Assert(IsInt(int64(3)), IsTrue)
}

func (s *CoreSuite) Test_IsFloat(c *C) {
	c.Assert(IsFloat(1), IsFalse)
	c.Assert(IsFloat(4.0), IsTrue)
	c.Assert(IsFloat(3.5), IsTrue)
	c.Assert(IsFloat(float32(3.5)), IsTrue)
}

func (s *CoreSuite) Test_MOD(c *C) {
	// c.Assert(MOD(4.0, 2.0), Equals, 0) // two floats
	// c.Assert(MOD(4, 2), Equals, 0)     // two ints
	// c.Assert(MOD(9.0, 3), Equals, 0)   // float, int
	c.Assert(MOD(7, 3.5), Equals, 1) // int, float
}

func (s *CoreSuite) Test_ADD(c *C) {
	// adding ints and floats
	c.Assert(ADD(3, 4.1), Equals, 7.1)
	c.Assert(ADD(4.1, 3), Equals, 7.1)
	// adding ints
	c.Assert(ADD(3, 4), Equals, 7.0)
	// NOTE ADD always returns a float
	// adding floats
	c.Assert(ADD(3.0, 4.1), Equals, 7.1)
	// c.Assert(func() { ADD("foo", 3) }, Panics, AddArgTypeError)
}
