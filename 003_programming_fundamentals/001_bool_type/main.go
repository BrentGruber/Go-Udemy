package main

import "fmt"

// Compiler assigns 0 value of false
var x bool

func main() {
	fmt.Println(x)
	x = true //change to true
	fmt.Println(x)

	a := 7
	b := 42

	// a == b is conditional expression that will evauluate to true or false
	fmt.Println(a == b)
}
