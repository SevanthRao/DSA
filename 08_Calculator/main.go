package main

import "fmt"

func main() {
	var n,m int
	fmt.Scanln(&n, &m)

	fmt.Println(n, "+", m,"=", n + m)
	fmt.Println(n, "-", m,"=", n - m)
	fmt.Println(n, "*", m,"=", n * m)
	fmt.Println(n, "/", m,"=", n / m)
	fmt.Println(n, "%", m,"=", n % m)
}