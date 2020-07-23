package piscine

import "github.com/01-edu/z01"

func IsNegative(n int) {
	if n < 0 {
		z01.PrintRune(84)
	} else {
		z01.PrintRune(70)
	}
	z01.PrintRune(10)
}
