package ctries_go

// ctries_go/ctrie_test.go

import (
	"code.google.com/p/intmath/intgr"
	// "encoding/binary"
	"fmt"
	xr "github.com/jddixon/xlattice_go/rnglib"
	. "launchpad.net/gocheck"
	"strings"
	// "sync/atomic"
	//"unsafe"
)

var _ = fmt.Print

func (s *XLSuite) dumpSlice(slice *[]byte) string {
	sl := *slice
	var ss []string
	for i := 0; i < len(sl); i++ {
		ss = append(ss, fmt.Sprintf("%02x ", sl[i]))
	}
	return strings.Join(ss, "")
}

func (s *XLSuite) insertHash(c *C, slice *[]byte, value byte) (where uint) {

	curSize := uint(len(*slice))
	c.Assert(where <= curSize, Equals, true)

	if curSize == 0 {
		*slice = append(*slice, value)
	} else {
		inserted := false
		var i uint
		var curValue, nextValue byte
		for i = 0; i < curSize-1; i++ {
			curValue = (*slice)[i]
			if curValue < value {
				nextValue = (*slice)[i+1]
				if nextValue < value {
					fmt.Printf("continuing: %02x after %02x, after %02x\n",
						value, curValue, nextValue)
					continue
				}
				fmt.Printf("inserting %02x after %02x, before %02x\n",
					value, curValue, nextValue)
				where = i + 1
				// do the insertion
				left := (*slice)[0:where]
				right := (*slice)[where:]
				fmt.Sprintf("%s + %02x + %s\n",
					s.dumpSlice(&left),
					value,
					s.dumpSlice(&right))
				left = append(left, value)
				left = append(left, right...)
				*slice = left
				inserted = true
				break
			}
		}
		if !inserted {
			fmt.Printf("%02x not inserted yet, examining i = %2d:  ",
				value, i)
			c.Assert(i, Equals, curSize-1)
			curValue = (*slice)[i]
			if curValue < value {
				fmt.Printf("appending %02x after %02x\n", value, curValue)
				*slice = append(*slice, value)
				where = curSize
			} else {
				fmt.Printf("inserting %02x before %02x\n", value, curValue)
				left := (*slice)[0:i]
				left = append(left, value)
				left = append(left, curValue)
				*slice = left
				where = curSize - 1
			}
		}
	}
	newSize := uint(len(*slice))
	c.Assert(newSize, Equals, curSize+1)
	fmt.Printf("  inserted 0x%02x at %d/%d\n", value, where, newSize)
	fmt.Printf("%s\n", s.dumpSlice(slice))
	return
}
func (s *XLSuite) TestInsert(c *C) {

	var (
		hc, bitmap      uint32
		lev             uint
		idx             uint
		flag, mask, pos int
		where           uint
	)
	rng := xr.MakeSimpleRNG()
	perm := rng.Perm(32) // a random permutation of [0..32)
	var slice []byte

	for i := byte(0); i < 32; i++ {
		hc = uint32(perm[i])
		// insert the value into the hash slice in such a way as
		// to maintain order
		idx = uint((hc >> lev) & 0x1f)
		c.Assert(idx, Equals, uint(hc))            // hc is restricted to that range
		where = s.insertHash(c, &slice, byte(idx)) // this is a uint

		flag = int(1 << (idx + 1))
		mask = flag - 1
		pos = intgr.BitCount(int(bitmap) & mask)

		occupied := uint32(1 << idx)
		// bitmap |= uint32(flag)
		bitmap |= uint32(occupied)

		fmt.Printf("%02d: hc %02x, idx %02x, mask 0x%08x, bitmap 0x%08x, pos %02d where %02d\n\n",
			i, hc, idx, mask, bitmap, pos, where)
	}

}
