package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	arr := make([]int, n)

	for i := 0; i <= n; i ++ {
		fmt.Scan(&arr[i])
	}

	fac := 1

	for _, p :=  range arr {
		for j := p; p >= 1; j--{
			fac = fac * j
		}
		if fac % 18 == 0 {
			fmt.Println(p)
		}
	}
}