package main

import (
	"strings"
	"testing"
)

func TestGenerate_num01(t *testing.T) {
	_num := Generate_num(1)
	num := strings.Split(_num, "")
	length := len(num)
	if length != 1 {
		t.Error("Test01 is failed")
	}
}

func TestGenerate_num02(t *testing.T) {
	_num := Generate_num(2)
	num := strings.Split(_num, "")
	length := len(num)
	if length != 2 {
		t.Error("Test02 is failed")
	} else if num[0] == num[1] {
		t.Error("Test02 is failed")
	}
}

func TestGenerate_num03(t *testing.T) {
	_num := Generate_num(3)
	num := strings.Split(_num, "")
	length := len(num)
	if length != 3 {
		t.Error("Test03 is failed")
	} else if num[0] == num[1] {
		t.Error("Test03 is failed")
	} else if num[1] == num[2] {
		t.Error("Test03 is failed")
	} else if num[2] == num[0] {
		t.Error("Test03 is failed")
	}
}

func TestEatBite01(t *testing.T) {
	eat, bite = EatBite(1, 1)
	if eat != 1 {
		t.Error("TestEatBite01 is failed")
	}
}
