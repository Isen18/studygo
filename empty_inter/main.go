package main

import "fmt"

func main() {
	print(nil)
	print(12)
	print("哈哈")

	print2(12)
	print2("哈哈哈")
}

func print(a interface{}) {
	// str, ok := a.(string)
	// if ok {
	// 	fmt.Printf("str=%s\n", str)
	// 	return
	// }

	// i, ok := a.(int)
	// if ok {
	// 	fmt.Printf("i=%d\n", i)
	// 	return
	// }

	if v, ok := a.(string); ok {
		fmt.Printf("string, v=%s\n", v)
	} else if v, ok := a.(int); ok {
		fmt.Printf("int, v=%d\n", v)
	}
}

func print2(a interface{}) {
	fmt.Printf("type: %T\n", a)

	// type switch
	switch v := a.(type) {
	case string:
		fmt.Println("string", v)
	case int:
		fmt.Println("int", v)
	}
}
