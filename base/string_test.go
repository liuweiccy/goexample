package base

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

func TestStringIntern(t *testing.T) {
	str1 := "Hello, World"
	str2 := "Hello, World"

	fmt.Printf("string addr:%p, %p\n", &str1, &str2)

	x1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	x2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("data addr: %v, %v\n", x1.Data, x2.Data)
}

func stringptr(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

type stringInterner map[string]string

func (si stringInterner) Intern(s string) string {
	if interned, ok := si[s]; ok {
		return interned
	}
	si[s] = s
	return s
}

func TestStringIntern2(t *testing.T) {
	si := stringInterner{}
	s1 := si.Intern("12")
	s2 := si.Intern(strconv.Itoa(12))

	fmt.Println(stringptr(s1) == stringptr(s2))
}
