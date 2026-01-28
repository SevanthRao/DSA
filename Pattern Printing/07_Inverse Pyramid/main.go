package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)

	for i:= n; i>= 1; i -- {
		for j:=1 ; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}