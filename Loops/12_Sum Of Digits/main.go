package main

import "fmt"

func main()  {
	var n int

	fmt.Scanln(&n)

	sum := 0

	for n!= 0{
		sum = sum + (n%10)
		n = n / 10
	}
	fmt.Println(sum)
}