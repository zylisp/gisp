package parser

import (
	"testing"

	. "github.com/masukomi/check"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) {
	TestingT(t)
}

type ParserSuite struct{}

var _ = Suite(&ParserSuite{})

func (s *ParserSuite) Test_exampleFn(c *C) {
	parsedNodes := ParseFromString("a-prog", "(def dbl (fn [x] (* 2 x)))")
	c.Assert(len(parsedNodes), Equals, 1)
	for _, node := range parsedNodes {
		c.Assert(NodeName(node.Type()), Equals, "NodeCall")
		c.Assert(node.String(), Equals, "(def dbl (fn [x] (* 2 x)))")
		// c.Assert(node, Equals, "NodeCall")
	}
}
