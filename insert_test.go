package ctries_go

// ctries_go/insert_test.go

/////////////////////////////////////////////////////////////////////
// THIS IS NOW LARGELY VESTIGIAL: this functionality is now in hamt_go
/////////////////////////////////////////////////////////////////////

import (
	"fmt"
	xr "github.com/jddixon/rnglib_go"
	xu "github.com/jddixon/xlUtil_go"
	. "gopkg.in/check.v1"
)

var _ = fmt.Print

//func (s *XLSuite) dumpSlice(slice *[]byte) string {
//	sl := *slice
//	var ss []string
//	for i := 0; i < len(sl); i++ {
//		ss = append(ss, fmt.Sprintf("%02x ", sl[i]))
//	}
//	return strings.Join(ss, "")
//}

func (s *XLSuite) insertHash(c *C, slice *[]byte, value byte) (where uint32) {

	curSize := uint32(len(*slice))
	c.Assert(where <= curSize, Equals, true)

	if curSize == 0 {
		*slice = append(*slice, value)
	} else {
		mySlice := *slice
		inserted := false
		var curValue, nextValue byte
		var i uint32
		for i = uint32(0); i < curSize-1; i++ {
			curValue = mySlice[i]
			if curValue < value {
				nextValue = mySlice[i+1]
				if nextValue < value {
					// fmt.Printf("continuing: %02x after %02x, after %02x\n",
					//	value, curValue, nextValue)
					continue
				}
				c.Assert(value < nextValue, Equals, true)
				where = i + 1
				//fmt.Printf("A: inserting %02x after %02x, before %02x, at %d\n",
				//	value, curValue, nextValue, where)
				// do the insertion
				var left []byte
				if where > 0 {
					left = append(left, mySlice[0:where]...)
				}
				right := mySlice[where:]
				//fmt.Printf("%s + %02x + %s => ",
				//	s.dumpSlice(&left),
				//	value,
				//	s.dumpSlice(&right))
				left = append(left, value)
				left = append(left, right...)

				//fmt.Printf("%s\n", s.dumpSlice(&left))
				*slice = left
				inserted = true
				break
			} else {
				c.Assert(value < curValue, Equals, true)
				where = i
				//fmt.Printf("B: inserting %02x before %02x at %d\n",
				//	value, curValue, where)
				// do the insertion
				var left []byte
				if where > 0 {
					left = append(left, mySlice[0:where]...)
				}
				right := mySlice[where:]
				// fmt.Printf("%s + %02x + %s\n",
				//	s.dumpSlice(&left), value, s.dumpSlice(&right))
				left = append(left, value)
				left = append(left, right...)
				*slice = left
				inserted = true
				break

			}
		}
		if !inserted {
			c.Assert(i, Equals, curSize-1)
			curValue = (*slice)[i]
			if curValue < value {
				//fmt.Printf("C: appending %02x after %02x\n", value, curValue)
				*slice = append(*slice, value)
				where = curSize
			} else {
				left := (*slice)[0:i]
				left = append(left, value)
				left = append(left, curValue)
				*slice = left
				where = curSize - 1
				//fmt.Printf("D: prepended %02x before %02x at %d\n",
				//	value, curValue, where)
			}
		}
	}
	newSize := uint32(len(*slice))
	c.Assert(newSize, Equals, curSize+1)
	//fmt.Printf("  inserted 0x%02x at %d/%d\n", value, where, newSize)
	// fmt.Printf("%s\n", s.dumpSlice(slice))
	return
}
func (s *XLSuite) TestInsert(c *C) {

	var (
		hc, bitmap      uint32
		lev             uint32
		idx             uint32
		flag, mask, pos uint32
		where           uint32
	)
	rng := xr.MakeSimpleRNG()
	perm := rng.Perm(32) // a random permutation of [0..32)
	var slice []byte

	for i := byte(0); i < 32; i++ {
		hc = uint32(perm[i])
		// insert the value into the hash slice in such a way as
		// to maintain order
		idx = (hc >> lev) & 0x1f
		c.Assert(idx, Equals, hc) // hc is restricted to that range
		where = s.insertHash(c, &slice, byte(idx))
		flag = uint32(1) << (idx + 1)
		mask = flag - 1
		//pos = intgr.BitCount(int(bitmap) & mask)
		pos = uint32(xu.BitCount32(bitmap & mask))
		occupied := uint32(1 << idx)
		bitmap |= occupied

		//fmt.Printf("%02d: hc %02x, idx %02x, mask 0x%08x, bitmap 0x%08x, pos %02d where %02d\n\n",
		//	i, hc, idx, mask, bitmap, pos, where)
		c.Assert(pos, Equals, where)
	}

}
