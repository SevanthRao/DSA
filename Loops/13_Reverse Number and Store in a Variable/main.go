package main

import "fmt"

func main()  {
	var n int64
	fmt.Scanln(&n)

	var rev int64 = 0

	for n !=0 {
		rev = (rev * 10) + (n % 10)
		n = n / 10
	}
	fmt.Println(rev)
}