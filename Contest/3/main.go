package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)


	lastDigit := 0
	// fmt.Println(n)


	for n != 0 {
		lastDigit = n % 10
		fmt.Println(lastDigit)
		n = n / 10
	}
}
