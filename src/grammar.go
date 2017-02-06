package main

import (
	"fmt"
)

func main() {
	//---------------------if------------
	var num int = 1
	//多分支
	if num < 0 {
		fmt.Println("num 负数")
	} else if num == 0 {
		fmt.Println("num 等于0")
	} else {
		fmt.Println("num 正数")
	}

	num = 0
	//多分支
	if num < 0 {
		fmt.Println("num 负数")
	} else if num == 0 {
		fmt.Println("num 等于0")
	} else {
		fmt.Println("num 正数")
	}

	num = -1
	//多分支
	if num < 0 {
		fmt.Println("num 负数")
	} else if num == 0 {
		fmt.Println("num 等于0")
	} else {
		fmt.Println("num 正数")
	}

	//---------------------switch------------

	var i = 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5, 6:
		fmt.Println("four, five, six")
	default:
		fmt.Println("invalid value!")
	}

	//经典的for语句 init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//精简的for语句 condition
	i = 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//死循环的for语句 相当于for(;;)
	i = 1
	for {
		if i > 10 {
			fmt.Println("break:", i)
			break
		}
		i++
	}

}
