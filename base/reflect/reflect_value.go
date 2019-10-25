package reflect

import (
	"fmt"
	"reflect"
	"strings"
)

func addr() {
	x := 2

	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()

	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())

	fmt.Println(&x)
}

func Print(x interface{})  {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type: %s\n", t.String())

	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("func (%s) %s%s\n", t, t.Method(i).Name, strings.TrimPrefix(methodType.String(), "func"))
	}
}
