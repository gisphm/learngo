package main

import "fmt"

func SumAndMultiply(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x, y := 3, 4
	xPLUSy, xTIMESy := SumAndMultiply(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
}
