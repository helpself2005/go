package main

import "time"
import "fmt"

func main() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C
	fmt.Println("timer expired!")
}
