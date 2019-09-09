package memo

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMemoGet(t *testing.T) {
	m := New(httpGetBody)

	incomingURLs := []string{
		"https://www.baidu.com/",
		"https://www.sourcetreeapp.com/",
		"https://im.qq.com/",
		"https://im.qq.com/",
		"https://www.sina.com.cn/",
	}

	var wg sync.WaitGroup
	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				fmt.Printf("发生错误%#v\n", err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}

	wg.Wait()
}
