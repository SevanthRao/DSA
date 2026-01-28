package main

import "fmt"

func main() {
	var letter string
	fmt.Scanln(&letter)

	vowels := [5]string{"a", "e", "i", "o", "u"}

	for _, char := range vowels {
		if char == letter {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
