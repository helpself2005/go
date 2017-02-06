package main

import (
	"fmt"
)

func main() {
	v, e := multi_ret("one")
	fmt.Println(v, e)
	//输出 1 true

	v, e = multi_ret("four")
	fmt.Println(v, e)
	//输出 0 false

	//通常的用法(注意分号后有e)
	if v, e = multi_ret("four"); e {
		fmt.Println("正常返回")
	} else {
		fmt.Println("错误返回")
	}

	sum(1, 2)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func multi_ret(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	var err bool
	var val int
	val, err = m[key]
	return val, err
}

func sum(nums ...int) {
	fmt.Print(nums, " ") //输出如 [1, 2, 3] 之类的数组
	total := 0
	for _, num := range nums { //要的是值而不是下标
		total += num
	}
	fmt.Println(total)
}
