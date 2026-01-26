package main

import "fmt"

func main()  {
	var n int

	fmt.Scanln(&n)

	i:= n

	for i >= 1 {
		fmt.Print(i, " ")
		i--
	}
}
