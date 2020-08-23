package main

import "code.google.com/p/go-tour/reader"

type MyReader struct{}

// define Read method to iterate buffer setting each element to 65 the number of ASCII A
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 65
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
