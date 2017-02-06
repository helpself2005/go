package main

import (
	"fmt"
)

func nextNum() func() int {

	i, j := 1, 1
	return func() int {

		var tmp = i + j
		i, j = j, tmp

		return tmp
	}
}

func fact(n int) int {

	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {

	nextNumFunc := nextNum()

	for i := 0; i < 10; i++ {
		fmt.Println(nextNumFunc())
	}

	fmt.Println("===========================================")

	fmt.Println(fact(5))

}
