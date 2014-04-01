package ctries_go

// ctries_go/ctrie.go

import (
	"sync/atomic"
	"unsafe"
)

type CTrie struct {
	root *INode
}

func NewCTrie() (ct *CTrie, err error) {

	// create a null INode
	i, err := NewINode(nil)

	if err != nil {
		ct = &CTrie{root: i}
	}
	return
}

// Returns a pointer to ct.root, the pointer having been read atomically.
func (ct *CTrie) READRoot() *INode {
	q := unsafe.Pointer(ct.root)
	p := atomic.LoadPointer(&q)
	r := (*INode)(p)
	return r
}	
func (ct *CTrie) insert (k KeyI, v interface{}) (err error) {

	r	:= ct.READRoot()
	if r == nil || r.IsNull() {
		var (
			m MainNodeI
			s	*SNode
			scn *CNode
			rn	*INode
		)
		s, err = NewSNode(k, v)
		if err == nil {
			m = MainNodeI(s)
			scn, err = NewCNode(m)
			if err == nil {
				rn, err = NewINode(scn)
				// XXX DO THE CAS
				_ = rn	// DEBUG
			}
		}

	} else { // test iinsert fails
		// XXX 
		// err = insert(k,v)
	}

	_ = r	// DEBUG
	return
}
