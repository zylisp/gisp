package core

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// hook up gocheck into the "go test" runner
type CoreSuite struct {
	suite.Suite
}

func TestCoreSuite(t *testing.T) {
	suite.Run(t, new(CoreSuite))
}
