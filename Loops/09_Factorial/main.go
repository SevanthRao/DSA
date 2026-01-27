package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)

	i:=n
	fact:=1

	for i >= 1 {
		fact = fact * i
		i--
	}
	fmt.Println(fact)
}