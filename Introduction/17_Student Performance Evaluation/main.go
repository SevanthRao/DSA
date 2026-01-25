package main

import "fmt"

func main()  {
	var Mark int64

	fmt.Scanln(&Mark)

	if Mark <= 100 && Mark > 90 {
		fmt.Println("Excellent")
	} else if Mark <= 90 && Mark > 80{
		fmt.Println("Good")
	} else if Mark <= 80 && Mark > 70 {
		fmt.Println("Fair")
	} else if Mark <= 70 && Mark > 60 {
		fmt.Println("Meets Expectations")
	} else {
		fmt.Println("Below Par")
	}
}