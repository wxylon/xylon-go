package xylon

import "testing"

func TestSum(t *testing.T) {
	v, e := Sum(2, 2)
	if v != 4 || e != nil {
		t.Errorf("Sum(2, 2) failed. Got %v, expected 4.", v)
	}
}
