package base

import "fmt"

func _() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)
}
