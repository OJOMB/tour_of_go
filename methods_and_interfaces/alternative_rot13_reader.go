package main

import (
	"io"
	"os"
	"strings"
)

func rot13(b byte) byte {
	var beg byte

	if b >= 'A' && b <= 'Z' {
		beg = 'A'
	} else if b >= 'a' && b <= 'z' {
		beg = 'a'
	} else {
		return b
	}

	return (((b - beg) + 13) % 26) + beg
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	size, err := r.r.Read(b)

	for i, v := range b {
		b[i] = rot13(v)
	}

	return size, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
