package main

import (
	"fmt"
	"runtime"
	"sync"
)

func f(i int) {
	defer wg.Done()
	fmt.Println(i)
	// wg.Done()
}

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	fmt.Println("num cpu", runtime.NumCPU())

	// for i := 0; i < 100; i++ {
	// 	wg.Add(1)
	// 	go f(i)
	// }

	wg.Wait()
	fmt.Println("main done")
}
