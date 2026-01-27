package main

import "fmt"

func main()  {
	var n int64
	fmt.Scanln(&n)

	if n == 0 {
		fmt.Println(0)
	}
	
	for n != 0 {
		fmt.Print(n % 10)
		n = n / 10
	}
}