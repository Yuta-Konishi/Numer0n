package main

import (
	"testing"
	"strings"
)

func Testgenerate_num01(t *testing.T) {
	num := generate_num(1)
	num = strings.Split(num, "")
	length := len(num)
	if length != 1 {
		t.Error("Test01 is failed")
	}
}

func Testgenerate_num02(t *testing.T) {
	num := generate_num(2)
	num = strings.Split(num, "")
	length := len(num)
	if length != 2 {
		t.Error("Test02 is failed")
	}else if num[0] == num[1] {
		t.Error("Test02 is failed")
	}
}

func Testgenerate_num03(t *testing.T) {
	num := generate_num(3)
	num = strings.Split(num, "")
	length := len(num)
	if length != 3 {
		t.Error("Test01 is failed")
	}else if num[0] == num[1] {
		t.Error("Test03 is failed")
	}else if num[1] == num[2] {
		t.Error("Test03 is failed")
	}else if num[2] == num[0] {
		t.Error("Test03 is failed")
	}
}
