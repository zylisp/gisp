package parser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NodeSuite struct {
	suite.Suite
}

func TestNodeSuite(t *testing.T) {
	suite.Run(t, new(NodeSuite))
}

func (s *NodeSuite) TestNodeName() {
	s.Equal(NodeName(0), "NodeIdent")
	s.Equal(NodeName(1), "NodeString")
	s.Equal(NodeName(2), "NodeNumber")
	s.Equal(NodeName(3), "NodeCall")
	s.Equal(NodeName(4), "NodeVector")
}
