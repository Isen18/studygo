package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// p := person{
	// 	Name: "张三",
	// 	Age:  int8(18),
	// }
	p := newPerson("张三", 18)

	//序列化
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Println(string(b))

	//反序列化
	str := `{"name":"张三","age":18}`
	var p2 person
	fmt.Printf("p2: %+v\n", p2)
	err = json.Unmarshal([]byte(str), &p2)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("p2: %+v\n", p2)
}

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int8   `json:"age"`
}

//构造函数
func newPerson(name string, age int8) *person {
	return &person{
		Name: name,
		Age:  age,
	}
}
