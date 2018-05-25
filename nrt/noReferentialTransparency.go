package nrt

/*
Result ...
*/
type Result struct {
	result int
}

func (r *Result) plus(a, b int) int {
	r.result = a + b
	return a + b
}

func mult(a, b int) int {
	return a * b
}
