package ctries_go

// ctries_go/sNode.go

type SNode struct {
	K	KeyI
	V	interface {}
	tomb bool
	MainNode
}

func NewSNode(k KeyI, v interface{}) (s *SNode, err error) {
	
	// XXX check parameters
	if err == nil {
		s = &SNode{		
			K : k,
			V : v,
			tomb: true,
		}
	}
	return
}
