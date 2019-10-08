package main

import "fmt"

var a, b = "Hello", "World"

var (
	d string = "111"
	f string
)

func main0() {
	c := 10
	fmt.Println("Hello, World!")
	fmt.Println(a, b, &c)
	c, g := 200, 300
	fmt.Println(&c, c)
	fmt.Println(g)
	fmt.Println(d, f)
}
