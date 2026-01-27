package main

import "fmt"

func main()  {
	var l, r int

	fmt.Scanln(&l, &r)

	for l <= r {
		fmt.Print(l, " ")
		l++
	}
	
}