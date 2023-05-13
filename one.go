package main

import "fmt"

const (
	a = iota + 10
	b
	c
	d
	e
	f
)

func main() {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("c = ", c)
	fmt.Println("d = ", d)
	fmt.Println("e = ", e)
	fmt.Println("f = ", f)

}
