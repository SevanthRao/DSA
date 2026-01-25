package main

import "fmt"

func main()  {
	var N int64
	var F int64

	fmt.Scanln(&N, &F)

	if N%F == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}