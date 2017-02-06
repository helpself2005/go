package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var total_tickets int32 = 10
var mutex = &sync.Mutex{}

func sell_tickets(i int) {

	for total_tickets > 0 {
		mutex.Lock()

		if total_tickets > 0 {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			total_tickets--
			fmt.Println("id: ", i, " ticket: ", total_tickets)
		} else {
			break
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		go sell_tickets(i)
	}

	var input string
	fmt.Scanln(&input)
	fmt.Println(total_tickets, "done")
}
