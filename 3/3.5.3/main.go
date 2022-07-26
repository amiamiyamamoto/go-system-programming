package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

func dumoChunc(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d byte)", string(buffer), length)
}
