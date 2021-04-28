package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		tmp := f1
		res := f1 + f2
		f1 = f2
		f2 = res
		return tmp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
