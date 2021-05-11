package main

import (
	"fmt"
	"unicode"
)

var (
	name string //名字
	age  int    //年龄
)

func main() {
	name = "go语言"
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

	for _, v := range name {
		// fmt.Println(v)
		fmt.Printf("%c\n", v)
	}

	// names := []byte(name)
	names := []rune(name)
	fmt.Println(string(names))

	s := "hello汉字"
	count := 0
	for _, c := range s {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Printf("count=%d\n", count)

	fmt.Println("Hello world!")
}
