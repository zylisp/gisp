package common

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// Hook up gocheck into the "go test" runner
type UtilSuite struct {
	suite.Suite
}

func TestUtilSuite(t *testing.T) {
	suite.Run(t, new(UtilSuite))
}

func (s *UtilSuite) TestRemoveExtension() {
	s.Equal(RemoveExtension("thing.zsp"), "thing")
	s.Equal(RemoveExtension("thing."), "thing")
	s.Equal(RemoveExtension("thing"), "thing")
	s.Equal(RemoveExtension("thingzsp"), "thingzsp")
}
