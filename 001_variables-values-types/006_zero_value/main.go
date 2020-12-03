package main

import "fmt"

// Default to zero value
var y string
var z int

func main() {

	//print y
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y) //print type

	//change y and print again
	y = "Bond, James Bond"
	fmt.Println(y)
	fmt.Printf("%T\n", y) // print type

	//print z
	fmt.Println("printing z:", z)
	fmt.Printf("%T\n", z) // print type
}
