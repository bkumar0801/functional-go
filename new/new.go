package new

/*
Iter ...
*/
type Iter chan interface{}

/*
New ...
*/
func New(elements ...interface{}) Iter {
	c := make(Iter)
	go func() {
		for _, element := range elements {
			c <- element
		}
		close(c)
	}()
	return c
}
