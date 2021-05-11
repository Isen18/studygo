package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:2000 failed, err=", err)
		return
	}

	fmt.Println("connected to 127.0.0.1:2000")
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read failed, err=", err)
			continue
		}

		text = strings.TrimSpace(text)
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println("write failed, err=", err)
			continue
		}
		fmt.Println("send msg to server, msg:", text)

		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read failed, err=", err)
			continue
		}

		fmt.Println("receive msg from server, msg:", string(buf[:n]))
		if string(buf[:n]) == "exit" {
			break
		}
	}

}
