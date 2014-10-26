package ctries_go

// ctries_go/bytesKey_test.go

import (
	// "encoding/binary"
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	. "gopkg.in/check.v1"
)

var _ = fmt.Print

func (s *XLSuite) TestBytesKey(c *C) {
	rng := xr.MakeSimpleRNG()

	for i := uint(0); i < 8; i++ {
		// length shouldn't matter, so long as it's > 8
		length := 8 + rng.Intn(32)
		data := make([]byte, length)
		rng.NextBytes(data)

		var expected uint32
		for j := uint(0); j < 8; j++ {
			// need to convert data[j] before shifting
			expected += uint32(data[j]) << (8 * j)
		}

		key := NewBytesKey(data)
		hc, err := key.Hashcode()
		c.Assert(err, IsNil)
		c.Assert(hc, Equals, expected)
	}

}
