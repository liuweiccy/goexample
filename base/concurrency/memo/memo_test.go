package memo

import (
	"fmt"
	"testing"
	"time"
)

func TestMemoGet(t *testing.T) {
	m := New(httpGetBody)

	incomingURLs := []string{
		"www.baidu.com",
		"www.facebook.com",
		"www.jd.com",
		"www.qq.com",
	}

	for _, url := range incomingURLs {
		start := time.Now()

		value, err := m.Get(url)
		if err != nil {
			fmt.Printf("发生错误%#v\n", err)
		}

		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
