package myfunc

import "testing"

func TestMyFunc(t *testing.T) {
	sum := Add(1, 2)

	if sum != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", sum)
	}
}
