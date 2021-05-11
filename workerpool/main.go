package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("worker-%d, start job:%d\n", id, i)
		results <- i * i
		fmt.Printf("worker-%d, finish job:%d\n", id, i)
	}
}

func main0() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	// 5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}

	close(jobs)

	//获取结果
	// for i := 0; i < 5; i++ {
	// 	res, ok := <-results
	// 	fmt.Println(res, ok)
	// }

	for i := range results {
		fmt.Printf("res: %d\n", i)
	}

	// for {
	// 	res, ok := <-results
	// 	fmt.Println(res, ok)
	// 	if !ok {
	// 		break
	// 	}
	// }
}

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("worker:%d start job:%d\n", id, j)
// 		// time.Sleep(time.Second)
// 		fmt.Printf("worker:%d end job:%d\n", id, j)
// 		results <- j * 2
// 	}
// }

// func main2() {
// 	jobs := make(chan int, 100)
// 	results := make(chan int, 100)
// 	// 开启3个goroutine
// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	// 5个任务
// 	for j := 1; j <= 5; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)
// 	// 输出结果
// 	for a := 1; a <= 5; a++ {
// 		res, ok := <-results
// 		fmt.Println(res, ok)
// 	}
// }

func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
