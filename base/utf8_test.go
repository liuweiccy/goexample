package base

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestUtf(t *testing.T) {
	b := make([]byte, utf8.UTFMax)

	n := utf8.EncodeRune(b, '好')
	fmt.Printf("%v -- %v\n", b, n)

	r, n := utf8.DecodeRune(b)
	fmt.Printf("%c -- %d\n", r, n)

	s := "大家好"
	for i := 0; i < len(s); {
		r, n := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%c:%v   ", r, n)
		i += n
	}
	fmt.Println()
	for i := len(s); i > 0; {
		r, n := utf8.DecodeLastRuneInString(s[:i])
		fmt.Printf("%c:%d    ", r, n)
		i -= n
	}
	fmt.Println()

	// 完整不一定有效，对于无效的字符串会转换为相应的无效的无效替代符('\uFFFD')，但是这个时候她又是完整的
	b = []byte("好")
	fmt.Printf("%t, ", utf8.FullRune(b))
	fmt.Printf("%t, ", utf8.FullRune(b[1:]))
	fmt.Printf("%t, ", utf8.FullRune(b[2:]))
	fmt.Printf("%t, ", utf8.FullRune(b[:1]))
	fmt.Printf("%t\n", utf8.FullRune(b[:2]))

	b = []byte("大家好")
	fmt.Println(utf8.RuneCount(b))
	fmt.Printf("%d, ", utf8.RuneLen('A'))
	fmt.Printf("%d, ", utf8.RuneLen('\u03a6'))
	fmt.Printf("%d, ", utf8.RuneLen('好'))
	fmt.Printf("%d, ", utf8.RuneLen('\U0010FFFF'))
	fmt.Printf("%d\n", utf8.RuneLen(0x1FFFFFFF))

	fmt.Printf("%t, ", utf8.RuneStart(b[0]))
	fmt.Printf("%t, ", utf8.RuneStart(b[1]))
	fmt.Printf("%t\n", utf8.RuneStart(b[2]))

	b = []byte("你好")
	fmt.Printf("%t,  ", utf8.Valid(b))
	fmt.Printf("%t,  ", utf8.Valid(b[1:]))
	fmt.Printf("%t,  ", utf8.Valid(b[2:]))
	fmt.Printf("%t,  ", utf8.Valid(b[3:]))
	fmt.Printf("%t,  ", utf8.Valid(b[:1]))
	fmt.Printf("%t,  ", utf8.Valid(b[:2]))
	fmt.Printf("%t\n", utf8.Valid(b[:3]))

	fmt.Printf("%t, ", utf8.ValidRune('好'))
	fmt.Printf("%t, ", utf8.ValidRune(0))
	fmt.Printf("%t, ", utf8.ValidRune(0xD800))
	fmt.Printf("%t\n", utf8.ValidRune(0x10FFFFFF))




}
