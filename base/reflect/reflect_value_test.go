package reflect

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestAddr(t *testing.T) {
	addr()
}

func TestAddr2(t *testing.T) {
	x := 2

	px1 := &x
	fmt.Printf("px1 = %v\n", px1)
	*px1 = 3
	fmt.Printf("x1 = %d\n", x)

	temp := reflect.ValueOf(&x).Elem()
	fmt.Printf("temp = %v\n", temp)
	px2 := temp.Addr().Interface().(*int)
	fmt.Printf("px2 = %v\n", px2)
	*px2 = 4
	fmt.Printf("x2 = %d\n", x)
}

func TestAddr3(t *testing.T) {
	// 通过内部方法设置
	x := 2
	temp := reflect.ValueOf(&x).Elem()
	temp.Set(reflect.ValueOf(4))
	fmt.Printf("x = %d\n", x)

	// 通过内部方法设置，特化的内部方法
	z := 2
	z1 := reflect.ValueOf(&z).Elem()
	z1.SetInt(4)
	fmt.Printf("z = %d\n", z)

	// 不同值类型的赋值，会panic
	m := 2
	m1 := reflect.ValueOf(&m).Elem()
	m1.Set(reflect.ValueOf(int64(8)))
	fmt.Printf("m = %d\n", m)

	// 对于不可寻址的值应用会panic
	y := 2
	y1 := reflect.ValueOf(&y)
	y1.Set(reflect.ValueOf(4))
	fmt.Printf("y = %d\n", y)
}

func TestAddr4(t *testing.T)  {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())

	fd := stdout.FieldByName("fd")
	fmt.Println(fd.CanAddr(), fd.CanSet())
}

func TestPrint(t *testing.T) {
	Print(time.Hour)
	Print(10)
}

func TestPrint2(t *testing.T) {
	fmt.Printf("%d %s\n", "hello", 42)
}