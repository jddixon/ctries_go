package ctries_go

const (
	VERSION      = "0.1.1"
	VERSION_DATE = "2014-04-06"
)

const (
	W          = uint(5)  // log base 2 of number of entries in a CNode
	LEVEL_MASK = 0x1f     // masks off W bits
	ARRAY_SIZE = uint(32) // exponentiation not a primitive in Go
)

const (
	OK = iota
	NOTFOUND
	RESTART
)
