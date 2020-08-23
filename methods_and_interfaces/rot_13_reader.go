package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(u uint8) uint8 {
	if u >= uint8(65) && u <= uint8(77) {
		u += uint8(13)
	} else if u >= uint8(77) && u <= uint8(90) {
		u -= uint8(13)
	} else if u >= uint8(97) && u <= uint8(109) {
		u += uint8(13)
	} else if u >= uint8(109) && u <= uint8(122) {
		u -= uint8(13)
	}
	return u
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	if err != nil {
		panic(err)
	}
	for index, item := range b {
		b[index] = rot13(item)
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number followed by return: ")
	r := rot13Reader{reader}
	_, _ = io.Copy(os.Stdout, &r)
	os.Stdin.Close()
}
