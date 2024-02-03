package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		message <- "hello"
	}()

	fmt.Println(<-message)
}
