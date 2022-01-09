package main

import "fmt"

const LoginToken string = "uidsl"  //pulic 

func main() {
	var username string = "Austin Amos"
	fmt.Println(username)
	fmt.Printf("Var is of type: %T \n", username)

	var isloggedIn bool = true
	fmt.Println(isloggedIn)
	fmt.Printf("Var is of type: %T \n", isloggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Var is of type: %T \n", smallVal)

	var smallfloat float32 = 255.467575855
	fmt.Println(smallfloat)
	fmt.Printf("Var is of type: %T \n", smallfloat)


	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Var is of type: %T \n", anotherVar)


	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberofUser := 300000.00
	fmt.Println(numberofUser)

	fmt.Println(LoginToken)
	fmt.Printf("Var is of type: %T \n", LoginToken)


}