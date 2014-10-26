package ctries_go

// ctries_go/sNode_test.go

import (
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
)

var _ = fmt.Print

// we are simple folk with simple needs
type DummyValue struct {
	n uint
}

func (s *XLSuite) TestSNode(c *C) {
	var err error
	rng := xr.MakeSimpleRNG()

	const COUNT = uint(8)
	keys := make([]BytesKey, COUNT)
	values := make([]interface{}, COUNT)
	hashCodes := make([]uint, COUNT)

	for i := uint(0); i < COUNT; i++ {
		length := 8 + rng.Intn(32) // so [8,40) bytes
		data := make([]byte, length)
		rng.NextBytes(data)
		key := NewBytesKey(data)
		var shc uint32
		var hc uint
		shc, err = key.Hashcode()
		hc = uint(shc)
		c.Assert(err, IsNil)
		keys[i] = *key
		hashCodes[i] = hc
		dummyVal := &DummyValue{i}
		values[i] = dummyVal
	}

	// a trivial test, presumably will be elaborated over time
	for i := uint(0); i < COUNT; i++ {
		var hc uint32
		hc, err = keys[i].Hashcode()
		c.Assert(uint(hc), Equals, hashCodes[i])
		expectedV := &DummyValue{n: i}
		c.Assert(values[i], DeepEquals, expectedV)
	}
}
