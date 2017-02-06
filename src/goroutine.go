package main

import (
	"fmt"
	"runtime"
	"time"
)

var a int = 1

func f(msg string) {
	fmt.Println(msg)
}

func sleep(i int) {
	for ; ; i += 1 {
		fmt.Println(i, "个屌丝")
		a += 1
	}
}

func main() {
	// go f("goroutine")

	// go func(msg string) {
	// 	fmt.Println(msg)
	// }("going")

	// time.Sleep(100)

	runtime.GOMAXPROCS(runtime.NumCPU())
	go sleep(1)
	time.Sleep(time.Microsecond)
	fmt.Println("end", 1)
}
