package closure

import "testing"

func TestAddx(t *testing.T) {
	add5 := Addx(5)
	add7 := Addx(7)
	got := add5(10)
	if got != 15 {
		t.Errorf("\nTest failed: \n\t\t expected= 15 \n\t\t actual= %d", got)
	}
	got = add7(10)
	if got != 17 {
		t.Errorf("\nTest failed: \n\t\t expected= 17 \n\t\t actual= %d", got)
	}
}
