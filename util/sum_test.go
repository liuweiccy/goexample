package util

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	if 3 == Sum(1, 2) {
		t.Log("success")
	} else {
		t.Fail()
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2)
	}
}

func TestSum2(t *testing.T) {
	myNum := []int{10, 20, 30, 40}
	sum := 0
	for _, num := range myNum {
		fmt.Printf("输出我的数字%d\n", num)
		sum += num
	}

	if sum != 100 {
		t.Error()
	}
}
