package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	names := os.Args
	runes := []rune(names[0])
	ind := 0
	length := 0

	for i, char := range runes {
		if char == '/' {
			ind = i
		}
		length = i + 1
	}

	for i := ind + 1; i < length; i++ {
		z01.PrintRune(runes[i])
	}
	z01.PrintRune(10)
}
