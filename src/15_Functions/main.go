package main

import "fmt"

func main() {

	fmt.Println("Lessons for functions")
	greeter() //excecutes the function.

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 3, 5, 7, 6, 8)
	fmt.Println("Pro result is: ", proRes, myMessage)

}

func greeter() {
	fmt.Println("Hello from Golang")
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Here is your pro adder result"
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}
