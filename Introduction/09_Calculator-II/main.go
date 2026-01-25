package main

import "fmt"

func main() {
	var n int64
	var m int64
	fmt.Scanln(&n)
	fmt.Scanln(&m)

	fmt.Printf("%d + %d = %d\n\n",n, m, n + m)
	fmt.Printf("%d - %d = %d\n\n",n, m, n - m)
	fmt.Printf("%d * %d = %d\n\n",n, m, n * m)
	fmt.Printf("%d / %d = %d\n\n",n, m, n / m)
	fmt.Printf("%d %% %d = %d\n\n",n, m, n % m)
}