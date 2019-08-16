package base

import (
	"fmt"
	"time"
)

func _() {
	request := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)
	limiter := time.Tick(time.Millisecond * 200)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}

	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request2", req, time.Now())
	}

}
