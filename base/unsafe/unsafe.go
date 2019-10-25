package unsafe

import (
	"unsafe"
)

func Print() {
}

func Float64Bits(f float64) uint64 {
	p := unsafe.Pointer(&f)
	return *(*uint64)(p)
}
