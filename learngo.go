package main

import (
	"fmt"
	"strconv"
)

<<<<<<< HEAD
type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an integer and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
=======
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
>>>>>>> 37e308e90d41dd863ce28b89c0a60cc17e8872c7
}
