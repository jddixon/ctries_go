package ctries_go

// ctries_go/bytesKey.go

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

var _ = fmt.Print

type BytesKey struct {
	slice []byte
}

func NewBytesKey(b []byte) *BytesKey {
	return &BytesKey{slice: b}
}

// KeyI interface ///////////////////////////////////////////////////

// convert the first 4 bytes of the key into an unsigned uint32
func (b *BytesKey) Hashcode() (hc uint32, err error) {
	buf := bytes.NewReader(b.slice)
	err = binary.Read(buf, binary.LittleEndian, &hc)
	if err != nil {
		fmt.Printf("attempt to read key failed\n")
	}
	return
}
