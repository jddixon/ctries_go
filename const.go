package ctries_go

const (
	VERSION			= "0.1.0"
	VERSION_DATE	= "2014-03-31"
)

const (
	W	= 8			// log base 2 of number of entries in a CNode
	ARRAY_SIZE = 32	// exponentiation not a primitive in Go 
)

const (
	OK	= iota
	NOTFOUND
	RESTART
)
