package main

import "fmt"

// https://pkg.go.dev/fmt

var y = 42

func main() {
	fmt.Println(y)         // print line (y\n)
	fmt.Printf("%T\n", y)  // print type
	fmt.Printf("%b\n", y)  // print binary
	fmt.Printf("%x\n", y)  // print in base 16
	fmt.Printf("%#x\n", y) // print in hex
	y = 911                // update
	fmt.Printf("%#x", y)   // print in hex

	fmt.Printf("%T\t%b\t%x\n", y, y, y) //Printf can take multiple args

	s := fmt.Sprintf("%T\t%b\t%x\n", y, y, y) //expression to print to string
	fmt.Println(s)

	fmt.Printf("%v\n", y)
}
