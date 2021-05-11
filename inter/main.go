package main

import "fmt"

func main() {
	c := cat{
		name: "cat1",
	}

	d := dog{
		name: "dog1",
	}

	speak(c)
	speak(d)

	var s speaker
	s = c
	s = d
	s.speak()

	s2 := speaker(c)
	s2.speak()

	s2 = d
	s2.speak()
}

type speaker interface {
	speak()
}

func speak(s speaker) {
	s.speak()
}

type cat struct {
	name string
}

//跟speaker有相同的方法，即实现了speaker接口
func (c cat) speak() {
	fmt.Printf("%s 喵喵喵\n", c.name)
}

type dog struct {
	name string
}

func (d dog) speak() {
	fmt.Printf("%s 汪汪汪\n", d.name)
}
