package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_num(n int) string {
	rand.Seed(time.Now().UnixNano())
	var x string
	var xi string
	x_list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < n; i++ {
		xi, x_list = delete(x_list, rand.Intn(10))
		x = x + xi
	}
	return x
}

func delete(slice []string, i int) (string, []string) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return "0", nil
	}
	ans := make([]string, len(slice))
	copy(ans, slice)

	ans = append(slice[:i], slice[(i+1):]...)

	return ret, ans
}

func main() {
	var n int
	fmt.Print("Enter the number of digits of the number to be generated : ")
	fmt.Scan(&n)
	x := generate_num(n)
	fmt.Println(x)
}
