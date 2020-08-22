package main

import "fmt"

// if we declare a const without explicit typing then it exists as an untyped const
// untyped consts will conform to types when required
// for example if we used Big in needFloat it will become float64

const (
	Big   = 1 << 100 // we can do this without an overflow error because numeric constants are high precision
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Printf("Big: %d", Big)
	fmt.Printf("Cast Big to float64: %f\n", float64(Big))

	fmt.Printf("needInt(Small): %d\n", needInt(Small))
	fmt.Printf("needFloat(Small): %d\n", needFloat(Small))
	bigNeedFloat := needFloat(Big)
	fmt.Println(bigNeedFloat) // we can do this because the scientific notation form used for floats brings 1<<100 into range
}
