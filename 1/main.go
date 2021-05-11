package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// hello()
	// println("main()", a)

	fmt.Println(strings.Split("abc", ""))
	fmt.Println(strings.Index("abc", ""))
}

var a string

func hello() {
	time.Sleep(time.Second)
	println("hello")
	a = "hello world"
	go f()
}

func f() {
	println("f")
	println("f()", a)
}
