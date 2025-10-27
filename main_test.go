package main

import "testing"

func TestMaxInt(t *testing.T) {
	a, b := 2, 7

	res := MaxInt(a, b)

	if res != b {
		t.Errorf("MaxInt(%d,%d)=%d; want %d", a, b, res, b)
	}
}

func TestMain(t *testing.M) {
	main()
}
