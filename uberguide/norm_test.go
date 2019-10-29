package uberguide

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestBase6(t *testing.T) {
	x := ""

	var (
		m []int
		n []int
	)

	if x == "" {
		m = []int{}
		n = nil
	}

	fmt.Println(len(m))
	fmt.Println(len(n))

	m = append(m, 1)
	n = append(n, 1)

	fmt.Println(m[0])
	fmt.Println(n[0])
}

// 尽量缩小变量的作用范围，除非在if外使用变量
func TestBase7(t *testing.T) {
	// Bad
	err := ioutil.WriteFile("test.txt", []byte("Hello"), 0644)
	if err != nil {
		fmt.Println(err)
	}

	// Good
	if err := ioutil.WriteFile("test.txt", []byte("Hello"), 0644); err != nil {
		fmt.Println(err)
	}
}

// 在if表达式外使用了变量，所以不放在if语句中，不能够减小变量的使用范围
func TestBase8(t *testing.T) {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}

	if err := decode(data); err != nil {
		return
	}
}

func decode(b []byte) error  {
	return nil
}
