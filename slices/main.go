package main

import (
	"fmt"
)

var (
	name = []string{"Bill", "Ted", "Frank"}
)

const (
	pi = 3.14
)

func main() {
	fmt.Printf("Hello, %s %v !\n", name[1:2], pi)
}
