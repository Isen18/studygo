package main

import (
	"fmt"
	"strconv"
	"sync"
)

var (
	x     = 0
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func add() {
	for i := 0; i < 50000; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
	wg.Done()
}

func main0() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println("x=", x)
}

var m sync.Map

// var m map[string]int

func main() {
	// m = make(map[string]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			fmt.Println(m.Load(key))

			// m[key] = n
			// fmt.Println(m[key])
			wg.Done()
		}(i)
	}

	wg.Wait()
}
