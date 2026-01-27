package main

import "fmt"

func main()  {
	var n int

	fmt.Scanln(&n)

	i := 1

	for i <= n {
		if i % 2 == 0 {
			fmt.Print(i, " ")
		}
		i++
	}
}