package main

import (
	"fmt"
)

type Skills []string

type Human struct {
	name   string
	age    int
	weight int
}

type Student struct {
	Human      //匿名字段, struct
	Skills     //匿名字段，自定义类型string slice
	int        //内置类型作为匿名字段
	speciality string
}

func main() {
	jane := Student{
		Human:      Human{"Jane", 35, 100},
		speciality: "Biology"}
	fmt.Println("Her name is ", jane.name)
	fmt.Println("Her age is ", jane.age)
	fmt.Println("Her weight is ", jane.weight)
	fmt.Println("Her specialityis ", jane.speciality)

	jane.Skills = []string{"anatomy"}
	fmt.Println("Her skills are ", jane.Skills)
	fmt.Println("She acquired two new ones ")
	jane.Skills = append(jane.Skills, "physics", "golang")
	fmt.Println("Her skills now are ", jane.Skills)

	jane.int = 3
	fmt.Println("Her preferred number is ", jane.int)
}
