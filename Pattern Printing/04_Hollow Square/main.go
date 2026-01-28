package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)

	for i:= 1; i<= n; i ++ {
		for j:=1 ; j <= n; j++ {
			if i == 1 || j == 1 || i == n || j == n{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}