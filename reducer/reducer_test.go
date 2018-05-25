package reducer

import (
	"testing"
)

func TestReduce(t *testing.T) {
	summer := func(memo interface{}, el interface{}) interface{} {
		return memo.(float64) + el.(float64)
	}
	if float64(.82)-Reduce(Float64(.1, .2, .3, .22), summer, float64(0)).(float64) > .000001 {
		t.Error("Sum Reduce failed")
	}
}
