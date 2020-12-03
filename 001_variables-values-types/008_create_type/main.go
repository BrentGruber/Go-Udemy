package main

import "fmt"

//declare a as type int
var a int

//create hotdog type with underlying int type
type hotdog int

//declare b as type hotdog
var b hotdog

func main() {
	a = 42
	b = 43

	//print a and type
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//print b and type
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	//cannot assign value of b to value of a because they are different types
	//a = b
}
