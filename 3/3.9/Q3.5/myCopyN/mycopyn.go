package mycopyn

import "io"

func CopyN(dest io.Writer, src io.Reader, length int) {
	buf := make([]byte, length)
	src.Read(buf)
	dest.Write(buf)
}
