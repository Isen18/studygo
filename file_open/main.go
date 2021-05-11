package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// readFromFile()
	// readFromBufio()
	readByUtils()
}

func readFromFile() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err=%v\n", err)
		return
	}
	defer file.Close()

	var buf [128]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, err=%v\n", err)
			return
		}

		fmt.Printf("读取了%d个字节\n", n)
		fmt.Print(string(buf[:n]))
	}
}

func readFromBufio() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err=%v\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("文件读完了")
			return
		}

		if err != nil {
			fmt.Printf("read from file failed, err=%v\n", err)
			return
		}

		fmt.Print(line)
	}
}

func readByUtils() {
	buf, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read from file failed, err=%v\n", err)
		return
	}

	fmt.Print(string(buf))
}
