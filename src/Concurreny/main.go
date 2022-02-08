package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}

}

func count(thing string, c chan string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
