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
		time.Sleep(time.Second * 1)

	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)

		}

	}

}
