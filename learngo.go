package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

func (h Human) String() string {
	return "<" + h.name + " - " + strconv.Itoa(h.age) + " years - ✆  " + h.phone + ">"
}

func main() {
	Bob := Human{"Bob", 39, "000-777-XXX"}
	fmt.Println("This Human is: ", Bob)
}
