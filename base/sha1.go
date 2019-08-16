package base

import (
	"crypto/sha1"
	"fmt"
)

func _() {
	s := "sha this string"

	h := sha1.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
