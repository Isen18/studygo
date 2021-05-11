package main

// func main() {
// 	slice := make([]string, 2, 4)
// 	Example(slice, "hello", 10)
// }

// func Example(slice []string, str string, i int) {
// 	debug.PrintStack()
// }

// type trace struct{}

// func main() {
// 	slice := make([]string, 2, 4)
// 	var t trace
// 	t.Example(slice, "hello", 10)
// }

// func (t *trace) Example(slice []string, str string, i int) {
// 	fmt.Printf("Receiver Address: %p\n", t)
// 	debug.PrintStack()
// }

func main() {
	var a [1]int
	c := a[:]
	// fmt.Println(c)
	println(c)
}
