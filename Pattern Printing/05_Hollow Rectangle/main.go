package main

import "fmt"

func main()  {
	var n, m int
	fmt.Scanln(&n, &m)

	for i:= 1; i<= n; i ++ {
		for j:=1 ; j <= m; j++ {
			if i == 1 || j == 1 || i == n || j == m{
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}