package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {

	for _, r := range s {
		c := rune(byte(r))
		z01.PrintRune(c)
	}
}
