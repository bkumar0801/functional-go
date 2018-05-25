package filter

import "github.com/functional-go/new"

/*
Predicate ...
*/
type Predicate func(interface{}) bool

/*
Uint64 ...
*/
func Uint64(elements ...uint64) new.Iter {
	c := make(new.Iter)
	go func() {
		for _, element := range elements {
			c <- element
		}
		close(c)
	}()
	return c
}

/*
Filter ...
*/
func Filter(pred Predicate, iter new.Iter) new.Iter {
	c := make(new.Iter)
	go func() {
		for element := range iter {
			if keep := pred(element); keep {
				c <- element
			}
		}
		close(c)
	}()
	return c
}
