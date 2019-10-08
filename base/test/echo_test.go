package test

import (
	"bytes"
	"fmt"
	"testing"
)
// 调用未导出的echo函数和改变全局变量out
// 测试的类型为白盒测试
func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, " ", []string{}, "\n"},
		{false, " ", []string{}, ""},
		{false, ":", []string{"1", "2", "3", "4", "5"}, "1:2:3:4:5"},
	}

	for _, test := range tests {
		desrc := fmt.Sprintf("echo(%v, %q, %q)\n", test.newline, test.sep, test.args)
		out = new(bytes.Buffer)

		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed:%v", desrc, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want:%q", desrc, got, test.want)
		}
	}
}
