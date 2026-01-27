package main

import "fmt"

func main()  {
	var n int

	fmt.Scanln(&n)
	
	i:=1
	for i <= 10 {
		fmt.Println(n, "*",i,"=",n*i)
		i++
	}
}