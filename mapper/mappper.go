package mapper

import (
	"github.com/functional-go/new"
)

/*
Mapper ...
*/
type Mapper func(interface{}) interface{}

/*
Map ...
*/
func Map(fn Mapper, iter new.Iter) new.Iter {
	c := make(new.Iter)
	go func() {
		for element := range iter {
			c <- fn(element)
		}
		close(c)
	}()
	return c
}
