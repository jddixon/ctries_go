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
func (ct *CTrie) READ_RootPtr() unsafe.Pointer {
	q := unsafe.Pointer(&ct.root)
	p := atomic.LoadPointer(&q)
	return p
}
func (ct *CTrie) insert(k KeyI, v interface{}) (err error) {

	older := ct.READ_RootPtr()
	r := (*INode)(older)
	if r == nil || r.IsNull() {
		var (
			m   MainNodeI
			s   *SNode
			scn *CNode
			rn  *INode
		)
		s, err = NewSNode(k, v)
		if err == nil {
			m = MainNodeI(s)
			scn, err = NewCNode(m, 0) // XXX level == 0 IS WRONG
			if err == nil {
				rn, err = NewINode(scn)
				if err == nil {
					q := unsafe.Pointer(&ct.root)
					newer := unsafe.Pointer(rn)
					ok := atomic.CompareAndSwapPointer(&q, older, newer)
					if !ok {
						err = ct.insert(k, v)
					}
				}
			}
		}

	} else { // test iinsert fails
		// XXX
		// err = insert(k,v)
	}

	_ = r // DEBUG
	return
}
