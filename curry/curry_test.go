package curry

import (
	"testing"
)

func TestCurry(t *testing.T) {
	plusOne := partialPlus(1)
	got := plusOne(5)
	if got != 6 {
		t.Errorf("\nTest failed: \n\t\t expected= 6 \n\t\t actual= %d", got)
	}
}
