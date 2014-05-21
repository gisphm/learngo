package main

import "fmt"

func myfunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func main() {
	x, y, z := 3, 4, 5
	myfunc(x, y, z, x+y, x*y, x+z)
}
