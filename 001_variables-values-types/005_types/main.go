package main

import "fmt"

// Go is a static programming language, variables are typed
// a dynamic programming language (javascript, python) does not require variables to hold types

//inferred int type
var y = 42

//inferred string type
var z = "Shaken, not stirred"

//raw string literal
//can include type even with inferred type from right side of =
var a string = `James said, "Shaken, not stirred"`

func main() {
	//println gives newline by default
	fmt.Println(y)

	//Formatted print, does not print new line by default
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)

	//can't do this
	//z = 43
	//because z is string type
}
