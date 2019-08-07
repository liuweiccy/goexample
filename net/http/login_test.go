package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"
)

// 访问总数
const N int = 100

// 并发数
const C int = 10

type config struct {
	url string
	total int
	clientNum int
}

func TestLogin(t *testing.T) {
	// todo 配置文件读取信息
	file, _ := os.Open("config.json")
	defer file.Close()

	decoder := json.NewDecoder(file)

	c := config{}
	decoder.Decode(&c)
	fmt.Println(c)

	group := &sync.WaitGroup{}
	group.Add(N)

	s := time.Now()
	for i := 0; i < C; i++ {
		go func() {
			for i := 0; i < N/C; i++ {
				testRequest(group)
			}
		}()
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
	//response, err := http.PostForm("http://192.168.20.97:8081/ws/logininfo", args)
	//if err != nil {
	//	fmt.Println("请求http地址发生错误：")
	//}
	// defer response.Body.Close()
	// body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))
}
