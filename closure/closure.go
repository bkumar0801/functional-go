package closure

/*
Addx ...
*/
func Addx(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
