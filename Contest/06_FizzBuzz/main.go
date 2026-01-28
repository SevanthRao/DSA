package main

import "fmt"

func main()  {
	var n int
	fmt.Scanln(&n)

	i:=1
	for i <= n {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {	
			fmt.Println("Buzz")
		} else if i % 3 == 0{
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}