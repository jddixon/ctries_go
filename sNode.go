package ctries_go

// ctries_go/sNode.go

type SNode struct {
	k    KeyI
	v    interface{}
	tomb bool // true
	MainNode
}

func NewSNode(k KeyI, v interface{}) (s *SNode, err error) {

	// XXX check parameters
	if err == nil {
		s = &SNode{
			k:    k,
			v:    v,
			tomb: true,
		}
	}
	return
}

// Not thread-safe.
func (sn *SNode) GetKey() KeyI {
	return sn.k
}

// Not thread-safe.
func (sn *SNode) GetValue() interface{} {
	return sn.v
}

// Not thread-safe.
func (sn *SNode) IsTomb() bool {
	return sn.tomb
}
