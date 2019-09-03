package typeassert

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestType1(t *testing.T) {
	var w io.Writer
	w = os.Stdout

	f := w.(*os.File)
	fmt.Println(f)

	if c, ok := w.(*bytes.Buffer); !ok {
		err := fmt.Errorf("类型断言失败:%v --> %v\n", reflect.TypeOf(c), reflect.TypeOf(w))
		fmt.Println(err)
	}
}

func TestFile(t *testing.T) {
	_, err := os.Open("c:/test.log")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
}

var ErrNotExist = errors.New("file does not exist")

func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

func TestFile2(t *testing.T) {
	_, err := os.Open("c:/test.log")
	fmt.Println(IsNotExist(err))
}

