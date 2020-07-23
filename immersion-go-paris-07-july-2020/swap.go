package piscine

func Swap(a *int, b *int) {

	var tmp int = 0
	tmp = *a
	*a = *b
	*b = tmp
}
