package main

import "fmt"


func main() {
	var length, breadth int

	fmt.Scanln(&length, &breadth)

	fmt.Println("Area =",length * breadth)
	fmt.Println("Perimeter =",2 * (length + breadth))
}
