package net

import (
	"fmt"
	"net"
)

func Client(port int)  {
	conn, err := net.Dial("tcp",":6985")
	defer conn.Close()
	if err!= nil {
		fmt.Println("链接错误:", err)
		return
	}

	conn.Write([]byte("test test test"))
}
