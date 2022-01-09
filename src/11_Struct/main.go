package main

import "fmt"

// structs => classes
func main() {

	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	austinUser := User{"Austine", "austine@go.dev", true, 16}
	fmt.Println(austinUser)
	fmt.Printf("austin details are: %+v\n", austinUser)
	fmt.Printf("Name is %v and email is %v.\n", austinUser.Name, austinUser.Email)

}

type User struct {
	Name   string
	Email  string
	Struct bool
	Age    int
}
