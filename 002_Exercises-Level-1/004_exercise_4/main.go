package main

import "fmt"

type poop int

var x poop

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
