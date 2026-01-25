package main

import "fmt"

func main()  {
	var Mark int64

	fmt.Scanln(&Mark)

	if Mark >= 35 && Mark <= 100 {
		fmt.Println("Pass")
	} else {
		fmt.Println("Fail")
	}
}