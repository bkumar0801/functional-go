package mapper

import (
	"reflect"
	"strings"
	"testing"

	"github.com/functional-go/new"
)

func TestMap(t *testing.T) {
	mapper := func(i interface{}) interface{} {
		return strings.ToUpper(i.(string))
	}
	got := Map(mapper, new.New("milu", "fekubaba"))
	item := new.New("MILU", "FEKUBABA")
	for x := range got {
		want := <-item
		if !reflect.DeepEqual(want, x) {
			t.Errorf("\nTest failed: \n\t\t expected= %v \n\t\t actual= %v", want, x)
		}
	}

}
