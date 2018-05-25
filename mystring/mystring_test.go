package mystring

import "testing"

func TestAppend(t *testing.T) {
	x := GetHello()
	y := Append(x, ", world")
	r1 := y
	r2 := y
	want := "Hello, world"
	if want != r1 {
		t.Errorf("\nTest failed 1: \n\t\t expected= %s \n\t\t actual= %s", want, r1)
	}

	if want != r2 {
		t.Errorf("\nTest failed 2: \n\t\t expected= %s \n\t\t actual= %s", want, r2)
	}

	//
	x = GetHello()
	r1 = Append(x, ", world")
	r2 = Append(x, ", world")
	if want != r1 {
		t.Errorf("\nTest failed 3: \n\t\t expected= %s \n\t\t actual= %s", want, r1)
	}

	if want != r2 {
		t.Errorf("\nTest failed 4: \n\t\t expected= %s \n\t\t actual= %s", want, r2)
	}
}
