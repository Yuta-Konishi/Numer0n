package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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
	x_arr := strings.Split(x, "")
	call_arr := strings.Split(call, "")

	for i := 0; i < len(call_arr); i++ {
		for j := 0; j < len(x_arr); j++ {
			if call_arr[i] == x_arr[j] {
				if i == j {
					eat++
				} else {
					bite++
				}
			}
		}
	}

	return eat, bite
}

func HighLow(x string) string {
	highlow := ""
	x_arr := strings.Split(x, "")

	for i := range x_arr {
		j, _ := strconv.Atoi(x_arr[i])
		if j <= 4 {
			highlow += "LOW"
		} else {
			highlow += "HIGH"
		}
		highlow += "*"
	}

	return highlow[:len(highlow)-1]
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
	var players int
	fmt.Print("Enter the number of players: ")
	fmt.Scan(&players)

	if players == 1 {
		var n, eat, bite int
		var call string

		fmt.Print("Enter the number of digits of the number to be generated: ")
		fmt.Scan(&n)
		x := Generate_num(n)

		//fmt.Println(x)

		for eat < n {
			fmt.Scan(&call)

			if call == "HIGH&LOW" {
				highlow := HighLow(x)
				fmt.Println(highlow)
			} else {
				eat, bite = EatBite(x, call)
				fmt.Printf("%dEAT %dBITE\n", eat, bite)
			}
		}

		fmt.Println("success!")

	} else if players == 2 {
		var x1, x2, call1, call2 string
		var eat1, bite1, eat2, bite2 int

		fmt.Print("Enter the number of player1: ")
		fmt.Scan(&x1)
		fmt.Print("Enter the number of player2: ")
		fmt.Scan(&x2)

		for {
			fmt.Print("Player1: ")
			fmt.Scan(&call1)
			if call1 == "HIGH&LOW" {
				highlow := HighLow(x2)
				fmt.Println(highlow)
			} else {
				eat1, bite1 = EatBite(x2, call1)
				fmt.Printf("%dEAT %dBITE\n", eat1, bite1)
			}

			fmt.Print("Player2: ")
			fmt.Scan(&call2)
			if call2 == "HIGH&LOW" {
				highlow := HighLow(x1)
				fmt.Println(highlow)
			} else {
				eat2, bite2 = EatBite(x1, call2)
				fmt.Printf("%dEAT %dBITE\n", eat2, bite2)
			}

			if eat1 == len(x2) && eat2 == len(x1) {
				fmt.Println("Draw")
				break
			} else if eat1 == len(x2) {
				fmt.Println("Player1 wins!")
				break
			} else if eat2 == len(x1) {
				fmt.Println("Player2 wins!")
				break
			}
		}
	} else {
		fmt.Println("Error: Only one or two players can participate.")
	}

}
