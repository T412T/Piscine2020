package piscine

func UltimateDivMod(a *int, b *int) {
	var tmpA = *a
	var tmpB = *b
	*a = *a / *b
	*b = tmpA % tmpB
}
