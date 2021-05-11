package main

import "fmt"

func main() {
	d := dog{
		animal: animal{name: "小黄"},
		feet:   4,
	}

	d.move()
	d.wang()

	fmt.Printf("%+v", d)
	fmt.Printf("name: %s", d.name)
}

type animal struct {
	name string
}

func (a *animal) move() {
	fmt.Println("move")
}

type dog struct {
	feet   int8
	animal //结构体的匿名嵌套
}

func (d *dog) wang() {
	fmt.Printf("%s wangwang~\n", d.name)
}
