package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64

	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := rect{2.9, 4.8}
	c := circle{4.3}

	//通过指针调用接口
	s := []shape{&r, &c}
	for _, sh := range s {
		fmt.Println(sh)
		fmt.Println(sh.area())
		fmt.Println(sh.perimeter())
	}

	sh1 := &r
	fmt.Println("接口指针调用：", sh1.area())

}
