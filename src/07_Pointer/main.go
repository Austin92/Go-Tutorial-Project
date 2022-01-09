package main

import "fmt"

func main() {

	// passing memory address. direct address to the value
	fmt.Println("Welcome to Pointers lesson")

	//var ptr *int
	//fmt.Println("Value of pointer is ", ptr)

	myNumber := 4

	//referencing using &
	var ptr = &myNumber

	fmt.Println("Value of the Pointer is ", ptr)
	fmt.Println("Value of the Pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is: ", myNumber)
}
