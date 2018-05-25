package nrt

import "testing"

func TestAdd(t *testing.T) {
	r := Result{}
	got := r.plus(2, mult(3, 4))
	want := 14
	if want != got {
		t.Errorf("\nTest failed 1: \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}
	if want != r.result {
		t.Errorf("\nTest failed 1: \n\t\t expected= %d \n\t\t actual= %d", want, r.result)
	}

	// replace expression mult(3, 4) with its return value
	r = Result{}
	got = r.plus(2, 12)
	if want != got {
		t.Errorf("\nTest failed 2: \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}
	if want != r.result {
		t.Errorf("\nTest failed 2: \n\t\t expected= %d \n\t\t actual= %d", want, r.result)
	}

	// replace expression r.plus(2, 12) with its return value
	r = Result{}
	got = 14
	if want != got {
		t.Errorf("\nTest failed 3: \n\t\t expected= %d \n\t\t actual= %d", want, got)
	}
	if want != r.result {
		t.Errorf("\nTest failed 3: \n\t\t expected= %d \n\t\t actual= %d", want, r.result)
	}
}
