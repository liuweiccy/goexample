package base

import (
	"fmt"
	"testing"
)

func TestByteCounterWrite(t *testing.T) {
	var c ByteCounter
	_, _ = c.Write([]byte("hello"))
	fmt.Println(c)
}

func TestByteCounterWrite2(t *testing.T) {
	var c ByteCounter
	var name = "Dolly"
	_, _ = fmt.Fprintf(&c, "Hello, %s", name)
	fmt.Println(c)
}

func TestByteCounterWrite3(t *testing.T) {
	var w WordCounter
	_, _ = w.Write([]byte("Hello Hey Hi,Ha"))
	fmt.Println(w)
}
