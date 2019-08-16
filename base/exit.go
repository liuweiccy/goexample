package base

import (
	"fmt"
	"os"
)

func _() {
	defer fmt.Println("!")
	os.Exit(3)
}
