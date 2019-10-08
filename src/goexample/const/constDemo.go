package main

import "fmt"

const Height = 10
const Width = 10

const (
	UnKnown = 0
	Female  = 1
	Male    = 2
)

func main() {
	area := Height * Width
	fmt.Println(area)
	fmt.Println(UnKnown)
}
