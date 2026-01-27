package main

import "fmt"

func main()  {
	var n int64
	var x int64
	fmt.Scanln(&x, &n)
	temp := x

	var i int64 = 1

	if n == 0{
		fmt.Println(1)
		return
	}

	for i < n {
		temp = temp * x
		i++
	}
	fmt.Println(temp)
}