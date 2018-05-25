package rt

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := plus(2, mult(3, 4))
	want := 14
	if want != got {
		t.Errorf("\nTest failed : \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}

	// replace expression mult(3, 4) with its return value
	got = plus(2, 12)
	if want != got {
		t.Errorf("\nTest failed : \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}

	// replace expression plus(2, 12) with its return value
	got = 14
	if want != got {
		t.Errorf("\nTest failed : \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}
}
