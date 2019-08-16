package base

import (
	"fmt"
	"os"
)

func _() {
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
