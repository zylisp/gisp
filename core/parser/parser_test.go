package parser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ParserSuite struct {
	suite.Suite
}

func TestParserSuite(t *testing.T) {
	suite.Run(t, new(ParserSuite))
}

func (s *ParserSuite) TestExampleFn() {
	parsedNodes := ParseFromString("a-prog", "(def dbl (fn [x] (* 2 x)))")
	s.Equal(len(parsedNodes), 1)
	for _, node := range parsedNodes {
		s.Equal(NodeName(node.Type()), "NodeCall")
		s.Equal(node.String(), "(def dbl (fn [x] (* 2 x)))")
		// s.Equal(node, "NodeCall")
	}
}
