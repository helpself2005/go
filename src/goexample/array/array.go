package main

import "fmt"

func main() {
	d := [...]int{1, 2, 3, 4, 5} // 根据元素的个数，设置数组的大小
	fmt.Println(len(d))          //[1 2 3 4 5]

	for i := 0; i < len(d); i++ {
		fmt.Printf("%d th element of is %d", i, d[i])
		fmt.Println("---")
	}
	sum := 0
	for i, v := range d {
		fmt.Printf("%d th range element of is %d", i, v)
		fmt.Println("---")
		sum += v
	}

	fmt.Println(sum)
}
