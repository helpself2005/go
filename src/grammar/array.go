package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("array a:", a)

	a[1] = 10
	a[3] = 30

	fmt.Println("assign:", a)
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("init:", b)

	var c [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}

	fmt.Println("2d: ", c)

	d := [5]int{1, 2, 3, 4, 5}

	e := d[2:4]
	// a[2] 和 a[3]，但不包括a[4]
	fmt.Println("e:", e)

	e = d[:4]
	// 从 a[0]到a[4]，但不包括a[4]
	fmt.Println("e:", e)

	e = d[2:]
	// 从 a[2]到a[4]，且包括a[2]
	fmt.Println("e:", e)
}
