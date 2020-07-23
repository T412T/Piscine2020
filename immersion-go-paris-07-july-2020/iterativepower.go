package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	s := 1
	for i := 1; i <= power; i++ {
		s = s * nb
	}
	return s
}
