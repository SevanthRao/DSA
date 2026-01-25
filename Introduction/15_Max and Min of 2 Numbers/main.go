package main

import "fmt"

func main()  {
	var A, B int64
	
	fmt.Scanln(&A, &B)

	if A < B {
		fmt.Println("Min =",A)
		fmt.Println("Max =",B)
	} else {
		fmt.Println("Min =",B)
		fmt.Println("Max =",A)
	}
}