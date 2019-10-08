package main

import (
	"fmt"
)

func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5] // slice point data pointer
	fmt.Println("array cd", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)

	var numbers = make([]int, 3, 5)

	printSlice(numbers)

	var numbers2 []int

	if numbers2 == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
