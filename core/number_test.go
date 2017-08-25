package core

import (
	. "github.com/masukomi/check"
)

type NumberSuite struct{}

var _ = Suite(&NumberSuite{})

func (s *NumberSuite) Test_AddNumbers(c *C) {
	// adding ints and floats
	sumNum, err := AddNumbers(Number{Value: 3}, Number{Value: 4.1})
	c.Assert(err, IsNil)
	c.Assert(sumNum, Equals, Number{Value: 7.1})

	sumNum, err = AddNumbers(Number{Value: 4.1}, Number{Value: 3})
	c.Assert(err, IsNil)
	c.Assert(sumNum, Equals, Number{Value: 7.1})

	sumNum, err = AddNumbers(Number{Value: 3}, Number{Value: 4})
	c.Assert(err, IsNil)
	c.Assert(sumNum, Equals, Number{Value: 7})

	sumNum, err = AddNumbers(Number{Value: 4.1}, Number{Value: 3.0})
	c.Assert(err, IsNil)
	c.Assert(sumNum, Equals, Number{Value: 7.1})
}

func (s *NumberSuite) Test_ToFloat(c *C) {
	x := Number{Value: 4}
	c.Assert(x.ToFloat(), Equals, 4.0)
	x.Value = 4.1
	c.Assert(x.ToFloat(), Equals, 4.1)
}

func (s *NumberSuite) Test_ToInt(c *C) {
	x := Number{Value: 4}
	c.Assert(x.ToInt(), Equals, 4)
	x.Value = 4.1
	c.Assert(x.ToInt(), Equals, 4)
}
