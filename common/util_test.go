package common

import (
	"testing"

	. "github.com/masukomi/check"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) {
	TestingT(t)
}

type CommonSuite struct{}

var _ = Suite(&CommonSuite{})

func (s *CommonSuite) Test_RemoveExtension(c *C) {
	c.Assert(RemoveExtension("thing.zsp"), Equals, "thing")
	c.Assert(RemoveExtension("thing."), Equals, "thing")
	c.Assert(RemoveExtension("thing"), Equals, "thing")
	c.Assert(RemoveExtension("thingzsp"), Equals, "thingzsp")
}
