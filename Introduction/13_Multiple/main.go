package main

import "fmt"

func main () {
	var M int64
	var N int64

	fmt.Scanln(&N, &M)

	if M%N == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}