package main

import "fmt"

func main()  {
	var n int

	fmt.Scanln(&n)

	i := 1
	for i <= n {
		fmt.Println(i)
		i++
	}
}
