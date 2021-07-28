package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate_num(n int) string {
	rand.Seed(time.Now().UnixNano())
	var x string
	var xi string
	x_list := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < n; i++ {
		xi, x_list = delete(x_list, rand.Intn(len(x_list)))
		x = x + xi
	}
	return x
}

func EatBite(x string, call string) (int, int) {
	eat, bite := 0, 0
	if x == call {
		eat = len(x)
	}

	return eat, bite
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
	var call string
	eat, bite := 0, 0
	fmt.Print("Enter the number of digits of the number to be generated : ")
	fmt.Scan(&n)
	x := Generate_num(n)
	fmt.Println(x)
	for eat < n {
		fmt.Scan(&call)
		eat, bite = EatBite(x, call)
		fmt.Printf("%dEAT %dBITE\n", eat, bite)
	}
}
