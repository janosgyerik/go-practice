package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(buf []byte) (int, error) {
	count := len(buf)
	for i := 0; i < count; i++ {
		buf[i] = 'A'
	}
	return count, nil
}

func main() {
	reader.Validate(MyReader{})
}
