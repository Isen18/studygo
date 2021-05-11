package main

import (
	"fmt"
	"net"
)

func processConn(conn net.Conn) {
	defer conn.Close()
	for {
		var b [512]byte
		n, err := conn.Read(b[:])
		if err != nil {
			fmt.Println("read failed, err=", err)
			continue
		}

		fmt.Println("receive msg from client, msg:", string(b[:n]))
		if string(b[:n]) == "exit" {
			break
		}

		_, err = conn.Write(b[:n])
		if err != nil {
			fmt.Println("write failed, err=", err)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("listen 127.0.0.1:2000 failed, err=", err)
		return
	}
	defer listener.Close()

	fmt.Println("listening on 127.0.0.1:2000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err=", err)
			return
		}

		go processConn(conn)
	}
}
