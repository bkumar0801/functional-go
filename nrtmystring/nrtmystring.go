package nrtmystring

/*
Append ...
*/
func Append(x *string, y string) string {
	*x += y
	return *x
}

/*
GetHello ...
*/
func GetHello() string {
	return "Hello"
}
