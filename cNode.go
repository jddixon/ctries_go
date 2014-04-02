package ctries_go

import (
	"fmt"
)

var _ = fmt.Print

type CNode struct {
	bmp   uint32 // bitmap for k == 5, so 2^k == 32
	Array [ARRAY_SIZE]MainNodeI
	MainNode
}

const KEY_MASK = uint(0x1f) // 5 bits

// Create a CNode, given a single MainNode.  If the MainNode is an
// SNode, we calculate the index appropriate for the level and
// insert the SNode in the CNode's Array in that slot,  If the MainNode
// is an INode, we thoughtlessly put it in slot 0 of the Array.  In
// either case we fill the other slots in the Array with null INodes.
//
func NewCNode(m MainNodeI, lvl uint) (c *CNode, err error) {

	// XXX check parameter

	var (
		isSNode bool
		ndx     uint
	)
	if err == nil {
		switch v := m.(type) {
		case *SNode:
			isSNode = true
			sn := m.(*SNode)
			ndx, err = sn.BitNdx(lvl)
		case *INode:
			// be happy
		default:
			fmt.Printf("NewCNode: MainNode of type %v\n", v)
			err = MainNodeUnknownType
		}
	}
	if err == nil {
		c = new(CNode)
		for i := uint(0); i < ARRAY_SIZE; i++ {
			if i == ndx && isSNode {
				c.bmp = uint32(1) << ndx
				// not a copy operation
				c.Array[i] = m
			} else {
				var nullINode *INode
				nullINode, err = NewINode(nil)
				if err != nil {
					break
				} else {
					c.Array[i] = nullINode
				}
			}
		}
	}
	return
}

// The index of the bit which will be ORed into the bitmap.
//
func (sn *SNode) BitNdx(lvl uint) (index uint, err error) {

	// XXX possible errors ignored

	if err == nil {
		key := sn.GetKey()
		hc := uint(key.Hashcode()) // cache me ?
		shiftCount := uint(W * lvl)

		index = ((KEY_MASK << (shiftCount)) & hc) >> shiftCount
	}
	return
}
