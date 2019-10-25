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
	// 16 （数据，长度）
	fmt.Println(unsafe.Sizeof("a"))
	// 8
	fmt.Println(unsafe.Sizeof(make(chan int)))
	fmt.Println(unsafe.Sizeof(make(map[string]string)))
	// 24 （数据，长度，容量）
	fmt.Println(unsafe.Sizeof(make([]int, 0)))

	// 8
	x := 1
	fmt.Println(unsafe.Sizeof(&x))

	// 8
	f := func() {}
	fmt.Println(unsafe.Sizeof(f))

	// 16 （类型，值）
	var i interface{} = 3
	fmt.Println(unsafe.Sizeof(i))
}

func TestPrint2(t *testing.T) {
	// 24
	fmt.Println(unsafe.Sizeof(struct {
		a bool
		b float64
		c int16
	}{}))

	// 16
	fmt.Println(unsafe.Sizeof(struct {
		a float64
		b int16
		c bool
	}{}))

	// 16
	fmt.Println(unsafe.Sizeof(struct {
		a float64
		b bool
		c int16
	}{}))

	// 1
	// 根据字段的最大的长度来进行对其（max:8）
	// 如 int16 对齐长度为2；int32 对其长度为4
	fmt.Println(unsafe.Alignof(true))

	// 8
	fmt.Println(unsafe.Sizeof(struct {
		c int32
		a int16
		b bool
	}{}))
}

func TestPrint3(t *testing.T) {
	var x = struct {
		a bool
		b int16
		c []int
	}{}

	// 32
	fmt.Println(unsafe.Sizeof(x))
	// 8
	fmt.Println(unsafe.Alignof(x))

	// x.a
	// 1,1,0
	fmt.Println(unsafe.Sizeof(x.a))
	fmt.Println(unsafe.Alignof(x.a))
	fmt.Println(unsafe.Offsetof(x.a))

	// x.b
	// 2,2,1
	fmt.Println(unsafe.Sizeof(x.b))
	fmt.Println(unsafe.Alignof(x.b))
	fmt.Println(unsafe.Offsetof(x.b))

	// x.a
	// 24,8,8
	fmt.Println(unsafe.Sizeof(x.c))
	fmt.Println(unsafe.Alignof(x.c))
	fmt.Println(unsafe.Offsetof(x.c))
}

func TestFloat64Bits(t *testing.T) {
	fmt.Printf("%#016x\n", Float64Bits(1.0))
}

func TestPrint4(t *testing.T) {
	var x struct{
		a bool
		b int16
		c []int
	}

	// 等价于 pb = &x.b
	pb := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42
	fmt.Println(x.b)

	// 下列语句块存在微妙的错误
	// 第一步是一个指针的数值，存在当发生GC时，存在内存块移动的情况，这时候地址会发生变化
	// 因为tmp仅仅是一个数值，不会随之改变，导致pb1指向的地址可能不是x.b的地址，导致赋值错误
	// 1
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	// 2
	pb1 := (*int16)(unsafe.Pointer(tmp))
	*pb1 = 34
	fmt.Println(x.b)
}
