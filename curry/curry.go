package curry

func plus(x, y int) int {
	return x + y
}
func partialPlus(x int) func(int) int {
	return func(y int) int {
		return plus(x, y)
	}
}
