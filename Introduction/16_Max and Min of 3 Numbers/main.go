package main

import "fmt"

func main() {
	var A, B, C int64

	fmt.Scanln(&A, &B, &C)

	if A <= B && B >= C && A <= C {
		fmt.Println("Min =", A)
		fmt.Println("Max =", B)
	} else if A <= C && C >= B && A <= B{
		fmt.Println("Min =", A)
		fmt.Println("Max =", C)
	} else if B <= A && A >= C && B <= C {
		fmt.Println("Min =", B)
		fmt.Println("Max =", A)
	} else if B <= C && C >= A && B <= A{
		fmt.Println("Min =", B)
		fmt.Println("Max =", C)
	} else if C <= A && A >= B && C <= B{
		fmt.Println("Min =", C)
		fmt.Println("Max =", A)
	} else {
		fmt.Println("Min =", C)
		fmt.Println("Max =", B)
	}
}
