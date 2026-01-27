package main

import "fmt"

func main()  {
	var n int64
	fmt.Scanln(&n)

	var rev int64 = 0

	var temp = n 

	for n != 0 {
		rev = (rev * 10) + (n % 10)
		n = n / 10
	}

	if rev == temp {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}