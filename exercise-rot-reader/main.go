package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const alphabetSize = byte('z' - 'a' + 1)

func (reader rot13Reader) Read(buf []byte) (int, error) {
	count, err := reader.r.Read(buf)
	for i := 0; i < count; i++ {
		buf[i] = rot13(buf[i])
	}
	return count, err
}

func rot13(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return rot13WithOffset(c, 'a')
	}
	if 'A' <= c && c <= 'Z' {
		return rot13WithOffset(c, 'A')
	}
	return c
}

func rot13WithOffset(c byte, offset byte) byte {
	return offset + (c + 13 - offset) % alphabetSize
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
