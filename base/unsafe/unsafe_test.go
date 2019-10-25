package unsafe

import (
	"fmt"
	"testing"
	"unsafe"
)

// http://golang-sizeof.tips/
// 可以查看数据的对齐方式以及填充字节的数量
func TestPrint(t *testing.T) {
	// 8
	fmt.Println(unsafe.Sizeof(int(0)))
	fmt.Println(unsafe.Sizeof(uint(0)))
	fmt.Println(unsafe.Sizeof(uintptr(0)))
	// 1
	fmt.Println(unsafe.Sizeof(int8(0)))
	// 1
	fmt.Println(unsafe.Sizeof(bool(true)))
	// 16
	fmt.Println(unsafe.Sizeof("a"))
	// 8
	fmt.Println(unsafe.Sizeof(make(chan int)))
	fmt.Println(unsafe.Sizeof(make(map[string]string)))
	// 24
	fmt.Println(unsafe.Sizeof(make([]int, 0)))

	// 8
	x := 1
	fmt.Println(unsafe.Sizeof(&x))

	// 8
	f := func() {}
	fmt.Println(unsafe.Sizeof(f))

	// 16
	var i interface{} = 3
	fmt.Println(unsafe.Sizeof(i))


}
