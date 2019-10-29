package performance

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkStrFmt(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

func BenchmarkStrConv(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}

func BenchmarkStrToByte(b *testing.B)  {
	w := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		w.Write([]byte("Hello World"))
	}
}

func BenchmarkStrToByte2(b *testing.B)  {
	w := new(bytes.Buffer)
	data := []byte("Hello World")
	for i := 0; i < b.N; i++ {
		w.Write(data)
	}
}
