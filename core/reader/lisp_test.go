package reader

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LispReaderSuite struct {
	suite.Suite
}

func TestLispReaderSuite(t *testing.T) {
	suite.Run(t, new(LispReaderSuite))
}
