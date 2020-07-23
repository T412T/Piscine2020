package piscine

func StrLen(s string) int {
	var c int = 0

	for _, r := range s {
		if byte(r) > 0 {
			c = c + 1
		}
	}

	return c
}
