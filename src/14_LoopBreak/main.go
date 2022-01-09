package main

import "fmt"

func main() {

	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	//for i := 0; i < len(days); i++ {
	//	fmt.Println(days[i])
	// }

	//for i := range days {
	//	fmt.Println(days[i])
	//}

	//for index, day := range days {

	//	fmt.Printf("index is %v and value is %v\n", index, day)

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")

}
