package ctries_go

// ctries_go/iNode.go

import (
//"sync/atomic"
//"unsafe"
)

type INode struct {
	main MainNodeI
}

func NewINode(m MainNodeI) (i *INode, err error) {
	i = &INode{main: m}
	return
}

func (i *INode) IsNull() bool {
	return i.main == nil
}
