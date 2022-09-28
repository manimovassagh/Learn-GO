package functions

func Multipling(x, y int) int {
	return x * y
}

func FuncLearnA(fn func(x, y int) int) int {
	return fn(12, 4)
}
