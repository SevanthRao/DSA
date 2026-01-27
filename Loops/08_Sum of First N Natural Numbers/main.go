package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)
	i:=1
	sum:=0
	for i <= n {
		sum = sum + i
		i++
	}
	fmt.Println(sum)
}