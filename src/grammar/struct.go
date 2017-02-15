package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	email string
}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2 * (r.height + r.width)
}

func main() {
	// person := Person{"Tom", 30, "tom@gmail.com"}
	// fmt.Println(person)

	// pPerson := &person
	// fmt.Println(pPerson)

	// pPerson.name = "shipl"
	// fmt.Println(pPerson)

	// person.email = "shipl@gmail.com"
	// fmt.Println(person)

	r := rect{width: 10, height: 15}
	fmt.Println("面积: ", r.area())
	fmt.Println("周长: ", r.perimeter())

	rp := &r
	fmt.Println("面积: ", rp.area())
	fmt.Println("周长: ", rp.perimeter())
}
