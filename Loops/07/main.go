package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	positive := 0
	negative := 0
	even := 0
	odd := 0

	for i := 1; i <= n; i++ {
		var inputN int
		fmt.Scan(&inputN)

		if inputN > 0 {
			positive++
		}
		if inputN < 0 {
			negative++
		}
		if inputN%2 == 0 {
			even++
		}
		if inputN %2 != 0 {
			odd++
		}
	}

	fmt.Println(positive)
	fmt.Println(negative)
	fmt.Println(even)
	fmt.Println(odd)
}
