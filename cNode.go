package ctries_go

type CNode struct {
	bmp	uint
	Array	[ARRAY_SIZE] MainNodeI
	MainNode
}

func NewCNode( m MainNodeI ) (c *CNode, err error) {

	// XXX check parameter

	c = &CNode{
		bmp:	uint(1),		// XXX
		}
	c.Array[0] = m

	return
}
