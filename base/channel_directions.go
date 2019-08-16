package base

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func _() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Passed Message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
