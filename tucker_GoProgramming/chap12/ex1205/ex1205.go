package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println("")

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a

	fmt.Println("")

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}

// [5]int 와 [10]int 는 다른 타입
// [5]int 와 [5]int64 는 다른 타입
