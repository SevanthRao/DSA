package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)

	count := 0
	if n == 0 {
		fmt.Println(1)
		return
	}

	for n != 0 {
		if n % 10 == 0 {
			count ++
		}
		n = n / 10
	}
	fmt.Println(count)
}