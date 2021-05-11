package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

func main() {
	// s := student{
	// 	Name: "张三",
	// 	Age:  18,
	// }

	// t := reflect.TypeOf(s)
	// v := reflect.ValueOf(s)
	// fmt.Printf("type:%s, kind:%s\n", t.Name(), t.Kind())
	// fmt.Printf("type:%s, value:%s\n", v.Type().Name(), v.FieldByName("Name").String())
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	fmt.Printf("name:%s, type:%s, jsonName:%s\n", field.Name, field.Type.Name(), field.Tag.Get("json"))
	// }

	var a int64 = 100
	// reflectSetValue1(a) //panic: reflect: reflect.Value.SetInt using unaddressable value
	reflectSetValue2(&a)
	fmt.Println(a)
}
