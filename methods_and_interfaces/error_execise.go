package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type TerribleError struct {
	when time.Time
	what string
}

func (e *TerribleError) Error() string {
	return fmt.Sprintf("error: %s, at: %v", e.what, e.when)
}

func Sqrt(x float64) (float64, error) {
	if x == 1 {
		return x, nil
	} else if x == 0 {
		return 0, nil
	} else if x < 0 {
		return 0, &TerribleError{time.Now().UTC(), "Input must be > 0"}
	}
	var acceptable_error float64 = 0.000000000000001
	var z float64 = 1

	for {
		old_z := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(old_z-z) < acceptable_error {
			return z, nil
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter number followed by return: ")
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := strconv.ParseFloat(
		strings.TrimSuffix(text, "\n"),
		64,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer, err := Sqrt(f)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(answer)

}
