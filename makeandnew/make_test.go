package makeandnew

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew1(t *testing.T) {
	i := new(int)

	var v int
	j := &v

	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))
}

func TestMake1(t *testing.T) {
	s := make([]int, 2, 10)
	m := make(map[string]int, 10)
	c := make(chan int, 5)

	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(m))
	fmt.Println(reflect.TypeOf(c))
}
