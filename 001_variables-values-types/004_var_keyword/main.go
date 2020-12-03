package main

import (
	"fmt"
)

//Best practice is to keep variable scope as narrow as possible

// must use var for global variables
// type is infered by right side of equal sign
var y = 43

// if no assignment, must assign type
// assigns zero value (default value) to variable
var z int

func main() {
	// short declaration operator
	x := 42
	fmt.Println(x)

	// var keyword
	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
