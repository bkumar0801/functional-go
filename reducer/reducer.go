package reducer

import "github.com/functional-go/new"

/*
Reducer ...
*/
type Reducer func(memo interface{}, element interface{}) interface{}

/*
Float64 ...
*/
func Float64(elements ...float64) new.Iter {
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
Reduce ...
*/
func Reduce(iter new.Iter, red Reducer, memo interface{}) interface{} {
	for element := range iter {
		memo = red(memo, element)
	}
	return memo
}
