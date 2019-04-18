package parser

import (
	. "github.com/masukomi/check"
)

type NodeSuite struct{}

var _ = Suite(&NodeSuite{})

func (s *NodeSuite) Test_nodeName(c *C) {
	c.Assert(NodeName(0), Equals, "NodeIdent")
	c.Assert(NodeName(1), Equals, "NodeString")
	c.Assert(NodeName(2), Equals, "NodeNumber")
	c.Assert(NodeName(3), Equals, "NodeCall")
	c.Assert(NodeName(4), Equals, "NodeVector")
}
