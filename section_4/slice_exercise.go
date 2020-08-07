package main

import "golang.org/x/tour/pic"

func kool(x, y int) uint8 {
	var rv int
	if y != 0 {
		rv = x * x / y * y
	} else {
		rv = 0
	}
	return rv % 255
}

func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		for j := 0; j < dy; j++ {
			s[i] = append(s[i], kool(i, j))
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
