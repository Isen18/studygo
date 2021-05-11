package main

import (
	"fmt"
)

func main() {
	// for i := 1; i < 10; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d*%d=%d\t", i, j, i*j)
	// 	}

	// 	fmt.Println()
	// }

	// s := []int{1, 2, 3}
	// s = append(s, 4)
	// fmt.Print(s)

	//字符统计
	// s := "hello hello world"
	// ss := strings.Split(s, " ")

	// m := make(map[string]int, len(ss))
	// // for _, w := range ss {
	// // 	m[w]++
	// // }

	// for _, w := range ss {
	// 	if _, ok := m[w]; !ok {
	// 		m[w] = 1
	// 	} else {
	// 		m[w]++
	// 	}
	// }

	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	//回文判断
	s := "a含b含ab"
	ss := []rune(s)
	len := len(ss)
	for i := 0; i < len>>1; i++ {
		if ss[i] != ss[len-1-i] {
			fmt.Println("不是回文")
			return
		}
	}

	fmt.Println("是回文")

}
