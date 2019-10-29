package uberguide

import "go.uber.org/atomic"

// 原子操作建议使用uber的原子包
type foo struct {
	count atomic.Uint32
}

func (f *foo)Add(num uint32) uint32 {
	f.count.Add(num)
	return f.count.Load()
}
