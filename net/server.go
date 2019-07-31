package net

import (
	"fmt"
	"net"
)

func Server1() {
	listener, err := net.Listen("tcp", ":6985")
	if err != nil {
		fmt.Println("listener error: ", err)
		return
	}

	fmt.Println("服务器开始启动开始监听端口：6985")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error: ", conn)
			return
		}

		go handlerConn(conn)
	}
}

func handlerConn(conn net.Conn) {
	defer conn.Close()

	for {
		b := make([]byte, 1024)
		conn.Read(b)
		fmt.Println("收到信息：", string(b[:]))
	}
}
