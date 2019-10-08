package main

import (
	"fmt"
)

func main() {
	name := "Hello World"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%d ", name[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
}
