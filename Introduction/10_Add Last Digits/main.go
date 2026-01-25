package main

import "fmt"

func main() {
	var n, m int

	fmt.Scanln(&n, &m)
	lastdigitN := n % 10
	lastdigitM := m % 10

	fmt.Println(lastdigitN + lastdigitM)
}