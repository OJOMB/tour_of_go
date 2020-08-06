package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux Baby")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Printf("text var is of type: %T\n", text)

	fmt.Printf("text[0] == 'o': %v\n", text[0] == 'o')

	switch {
	case strings.HasPrefix(text, "o"):
		fmt.Println("ahh yes, input that starts with an o")
	case strings.HasPrefix(text, "e"):
		fmt.Println("oh no, input should never start with an e!!!")
	default:
		fmt.Println("all input should start with an o dont you know?")
	}

}
