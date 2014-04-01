package ctries_go

// ctries_go/iNode.go

import (
	//"sync/atomic"
	//"unsafe"
)

type INode struct {
	main	MainNodeI
}

func NewINode (m MainNodeI) (i *INode, err error) {
	i = &INode{main: m}
	return
}

func (i *INode) IsNull() bool {
	return i.main == nil
}
	
// Returns a pointer to INode.main, the pointer having been read atomically.
//func (i *INode) READ() *MainNode {
//	q := unsafe.Pointer(i.main)
//	p := atomic.LoadPointer(&q)
//	r := (*MainNode)(p)
//	return r
//}	
