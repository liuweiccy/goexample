package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {
	param := "{\"clientType\":\"1\",\"account\":\"4735744240362690\",\"type\":\"1\",\"token\":\"66684A8B3733498AA5289D068DA96637\"}"

	response, err := http.Post("http://192.168.20.97:8081/ws/logininfo", "application/json;charset=utf-8", bytes.NewBuffer([]byte(param)))
	if err != nil {
		t.Error("请求http地址发生错误：")
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	t.Log(string(body))
}

func BenchmarkLogin(b *testing.B) {
	response, err := http.Get("http://192.168.20.97:8081/ws/logininfo")
	if err != nil {
		b.Error("请求http地址发生错误：")
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	b.Log(string(body))
}
