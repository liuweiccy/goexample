package uberguide

import (
	"fmt"
	"sync"
	"testing"
)

type S struct {
	data string
}

// 值接收器
// 允许值调用和指针调用
func (s S) Read() string {
	return s.data
}

// 指针接收器
// 只允许指针调用
func (s *S) Write(str string) {
	s.data = str
}

func TestBase(t *testing.T) {
	sVal := map[int]S{1: {"A"}}
	sVal[1].Read()
	// 编译不通过，指针接收器只能通过指针调用
	// sVal[1].Write("test")

	sPtr := map[int]*S{1: {"A"}}
	// 指针调用，可以调用值接收器和指针接收器方法
	sPtr[1].Read()
	sPtr[1].Write("Test")
}

type F interface {
	f()
}

type S1 struct{}

func (s S1) f() {}

type S2 struct{}

func (s *S2) f() {}

func TestBase2(t *testing.T) {
	s1Val := S1{}
	s1Ptr := &S1{}
	s2Val := S2{}
	s2Ptr := &S2{}

	var i F
	i = s1Val
	i = s1Ptr
	i = s2Ptr
	// s2Val 是一个值，不能够调用其指针接收器
	// i = s2Val
	i.f()
	// 为什么一个值能够调用，指针接收器的方法？
	// 是因为会制动转换为(&s2Val).f()
	s2Val.f()
}

// 私有类型采用类型嵌入
type smap struct {
	sync.Mutex
	data map[string]string
}

func newSMap() *smap {
	return &smap{
		data: make(map[string]string),
	}
}

func (m *smap) Get(k string) string {
	m.Lock()
	defer m.Unlock()

	return m.data[k]
}

// 导出类型采用专用字段
type SMap struct {
	mu   sync.Mutex
	data map[string]string
}

func NewSMap() *SMap {
	return &SMap{
		data: make(map[string]string),
	}
}

func (m *SMap) Get(k string) string {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.data[k]
}

// slice和map作为参数传入是指针应用的，你要确定你是否对他进行修改
// 同理当我们在返回，map和slice时，也不应该暴露出我们的内部状态
type Trip string
type Driver struct {
	trips []Trip
}

// 外部修改trips，会影响到d.trips的值，因为指向的地址相同
func (d *Driver) setTrips1(trips []Trip) {
	d.trips = trips
}

// 外部修改trips，不会影响到d.trips的值
func (d *Driver) setTrips2(trips []Trip) {
	d.trips = make([]Trip, len(trips))
	copy(d.trips, trips)
}

func TestBase3(t *testing.T) {
	trips := []Trip{"Chengdu", "Beijing", "Shanghai"}
	d1 := Driver{}
	d1.setTrips1(trips)
	trips[0] = "Chongqing"

	fmt.Println(d1.trips)
	fmt.Println(trips)
}

func TestBase4(t *testing.T) {
	trips := []Trip{"Chengdu", "Beijing", "Shanghai"}
	d1 := Driver{}
	d1.setTrips2(trips)
	trips[0] = "Chongqing"

	fmt.Println(d1.trips)
	fmt.Println(trips)
}

func TestBase5(t *testing.T) {
	ch := make(chan int)

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch)

	ch1 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 1

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
}

type Ops int

// Add = 1, Sub = 2, Mul = 3
const (
	Add Ops   = iota + 1
	Sub
	Mul
)
