package main

import (
	"os"
	"fmt"
)

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
