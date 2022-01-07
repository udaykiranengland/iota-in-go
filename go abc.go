package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(a, c, b)
}
