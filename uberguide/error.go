package uberguide

import (
	"errors"
	"fmt"
)

// 当需要检测错误时，且不需要额外的信息时，采用errors.New的方式创建错误
var ErrCouldNotOpen = errors.New("could not open")

func Open() error {
	return ErrCouldNotOpen
}

func use1() {
	if err := Open(); err != nil {
		if err == ErrCouldNotOpen {
			fmt.Println("处理错误：", err)
		} else {
			panic("Unknown Error")
		}
	}
}

// 当需要检测错误时，且不需要额外的信息时，采用自定义实现Error()方法来创建错误信息
type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %s is not found\n", e.file)
}

func open(file string) error {
	return errNotFound{file: file}
}

func use2() {
	if err := open("data.txt"); err != nil {
		if _, ok := err.(errNotFound); ok {
			fmt.Println("错误处理：", err)
		} else {
			panic("Unknown Error")
		}
	}
}

// 自定义的错误导出要注意，以为他成为公共API的一部分了
// 建议采用以提供接口的方式，来进行进行错误的检查
func IsNotFoundError(err error) bool {
	_, ok := err.(errNotFound)
	return ok
}

func use3() {
	if err := open("data.txt"); err != nil {
		if IsNotFoundError(err) {
			fmt.Println("错误处理：", err)
		} else {
			panic("Unknown Error")
		}
	}
}
