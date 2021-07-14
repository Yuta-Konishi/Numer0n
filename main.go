package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate_num(n int) int {
	rand.Seed(time.Now().UnixNano())
	x_list := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	x, x_list := delete(x_list, rand.Intn(10))
	return x
}

func delete(slice []int, i int) (int, []int) {
	ret := slice[i]
	if len(slice) < i || len(slice) < i {
		return 0, nil
	}
	ans := make([]int, len(slice))
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
