package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// initialise Reader
	var r io.Reader = strings.NewReader("Hello, Reader!")

	fmt.Println("============readInLoop===============")
	// initialise buffer as element byte slice
	b := make([]byte, 32)
	readInLoop(r, b)
	fmt.Println(b)

	s := string(b)
	fmt.Println(s)

	fmt.Println("===========readToBuffer==============")
	r = strings.NewReader("Hello, Reader!")
	var buf *bytes.Buffer = readToBuffer(r)
	fmt.Println(*buf)

}

func readInLoop(r io.Reader, b []byte) (n int, err error) {
	n = 0
	for {
		bytes_read, err := r.Read(b[n:])
		fmt.Printf("bytes read: %d\n", bytes_read)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		n += bytes_read
	}
	return n, nil
}

func readToBuffer(stream io.Reader) *bytes.Buffer {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf
}
