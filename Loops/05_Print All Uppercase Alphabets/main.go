package main

import "fmt"

func main()  {
	alpha := 'A'
	for alpha <= 'Z' {
		fmt.Printf("%c ", alpha)
		alpha++
	}
}