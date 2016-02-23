package ctries_go

// ctries_go/atomicity_test.go

import (
	"fmt"
	. "gopkg.in/check.v1"
	// "sync/atomic"
	"unsafe"
)

var _ = fmt.Print

func (s *XLSuite) TestBasics(c *C) {

	ct, err := NewCTrie()
	c.Assert(err, IsNil)
	c.Assert(ct, NotNil)

	rootPtr := &ct.root
	mainPtr := &ct.root.main

	p := ct.READ_RootPtr()

	fmt.Printf("ct (ptr)  = %v\n", unsafe.Pointer(ct))
	fmt.Printf("rootPtr   = %v\n", rootPtr)
	fmt.Printf("mainPtr   = %v\n", mainPtr)

	// initially root.main is set to nil
	c.Assert(p, Equals, unsafe.Pointer(mainPtr))
	q := *mainPtr
	c.Assert(q, IsNil)
}
