package main

import "fmt"

func main()  {
	var x, y int64

	fmt.Scanln(&x, &y)

	if x == 0 && y == 0 {
		fmt.Println("Origin")
	} else if x != 0 && y == 0 {
		fmt.Println("X axis")		
	} else if x == 0 && y != 0 {
		fmt.Println("Y axis")
	} else if x > 0 && y > 0 {
		fmt.Println("1st Quadrant")
	} else if x < 0 && y > 0 {
		fmt.Println("2nd Quadrant")
	} else if x < 0 && y < 0  {
		fmt.Println("3rd Quadrant")
	} else{
		fmt.Println("4th Quadrant")
	}
}