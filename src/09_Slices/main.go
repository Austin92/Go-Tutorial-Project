package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices (Memory!!)")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 265
	highScores[1] = 965
	highScores[2] = 465
	highScores[3] = 777
	// highScore[4] - 789

	highScores = append(highScores, 565, 878, 667)

	fmt.Println(highScores)

	//true
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//true
	fmt.Println(sort.IntsAreSorted(highScores))


	//how to remove a value from alices based on index

	var courses = []string{"Math", "English", "Physics", "History"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
