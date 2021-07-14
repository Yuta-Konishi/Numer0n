package main

import "testing"

func Testgenerate_num01(t *testing.T) {
	num := generate_num(1)
	if (num - 10) >= 0 {
		t.Error("Test01 is failed")
	}
}
