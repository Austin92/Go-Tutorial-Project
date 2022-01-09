package main

import "fmt"

func main() {

	fmt.Println("Welcome to array lesson")

	var fruitList [3]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomatos"
	fruitList[2] = "Orange"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Fruit List is: ", len(fruitList))

	var vegList = [3]string{"eggs", "beans", "rice"}
	fmt.Println("veg list is: ", vegList)

}
