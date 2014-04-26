package ctries_go

import (
	// . "launchpad.net/gocheck"
	"gopkg.in/check.v1"
	"testing"
)

// IF USING gocheck, need a file like this in each package=directory.

func Test(t *testing.T) { check.TestingT(t) }

type XLSuite struct{}

var _ = check.Suite(&XLSuite{})
