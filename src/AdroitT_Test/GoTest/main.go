package main

import (
	"fmt"
)

//run test or debug test or unit testing
func Calculation(x int) (result int) {
	result = x + 2
	return result
}

func main() {
	fmt.Println("Go Testing")
	result := Calculation(2)
	fmt.Println(result)
}
