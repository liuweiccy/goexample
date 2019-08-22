package strings

import (
	"fmt"
	"testing"
)

func TestNewReader(t *testing.T) {
	s := "hello"
	reader := NewReader(s)
	b := make([]byte, len(s))
	n, _ := reader.Read(b)
	fmt.Printf("读取到的字节数：%d,字符串：%s\n", n, string(b))
}

func TestLimitReader(t *testing.T) {
	s := "hello"
	reader := NewReader(s)
	r := LimitReader(reader, 5)
	b1 := make([]byte, 3)
	b2 := make([]byte, 3)
	n, err := r.Read(b1)
	fmt.Printf("b1读取到的字节数：%d,字符串：%s,错误：%v\n", n, string(b1), err)
	n, err = r.Read(b2)
	fmt.Printf("b2读取到的字节数：%d,字符串：%s,错误：%v\n", n, string(b1), err)
}
