package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(p []byte) (n int, err error) {
	n , err = len(p), nil
	for i := range p {
		p[i] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
