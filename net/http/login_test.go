package http

import (
	"fmt"
	"github.com/panjf2000/ants"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"
)

// 访问总数
const N int = 1000

// 并发数
const C int = 10000

func TestLogin(t *testing.T) {
	defer ants.Release()
	pool, _ := ants.NewPool(C)

	group := &sync.WaitGroup{}
	group.Add(N)

	s := time.Now()
	for i := 0; i < N; i++ {
		_ = pool.Submit(func() {
			testRequest(group)
		})
	}

	group.Wait()

	fmt.Printf("请求%d次总共花费时间：%d ms, QPS:%f \n", N, time.Now().Sub(s).Milliseconds(), float64(N)/time.Now().Sub(s).Seconds())
}

func testRequest(group *sync.WaitGroup) {
	args := url.Values{}
	args.Set("clientType", "1")
	args.Set("account", "4735744240362690")
	args.Set("type", "1")
	args.Set("token", "66684A8B3733498AA5289D068DA96637")
	_, _ = http.PostForm("http://192.168.20.97:8081/ws/logininfo", args)
	group.Done()
}
