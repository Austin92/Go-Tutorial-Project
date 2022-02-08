package main

import (
	"fmt"
	"time"
)

func main() {
	//multiple channels
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		c2 <- "Every two seconds"
		time.Sleep(time.Second * 2)

	}()

	for {
		fmt.Println(<-c1)
		fmt.Println(<-c2)

	}

}
