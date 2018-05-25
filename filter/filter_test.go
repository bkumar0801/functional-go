package filter

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	pred := func(i interface{}) bool {
		return i.(uint64) > 5
	}
	got := Filter(pred, Uint64(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	elements := Uint64(6, 7, 8, 9, 10)
	for x := range got {
		want := <-elements
		if !reflect.DeepEqual(want, x) {
			t.Errorf("\nTest failed: \n\t\t expected= %v \n\t\t actual= %v", want, x)
		}
	}
}
