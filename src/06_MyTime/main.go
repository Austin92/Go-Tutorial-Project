package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// to get the correct date and time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create date
	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}
