package core

import (
	. "github.com/masukomi/check"
	"testing"
)

// hook up gocheck into the "go test" runner
func Test(t *testing.T) { TestingT(t) }

type CoreSuite struct{}

var _ = Suite(&CoreSuite{})

func (s *CoreSuite) Test_isInt(c *C) {
	c.Assert(isInt(1), IsTrue)
	c.Assert(isInt(4.0), IsFalse)
}

func (s *CoreSuite) Test_isFloat(c *C) {
	c.Assert(isFloat(1), IsFalse)
	c.Assert(isFloat(4.0), IsTrue)
	c.Assert(isFloat(3.5), IsTrue)
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
	// adding floats
	c.Assert(ADD(3.0, 4.1), Equals, 7.1)
}
