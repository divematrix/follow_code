package main

import "fmt"

const Y int = 3

func main() {
	// x := 5
	// a := [x]int{1, 2, 3, 4, 5} error 개수에 변수를 쓸 수 없다

	b := [Y]int{1, 2, 3}

	var c [10]int

	fmt.Println(b[1])
	fmt.Println(c[0])
}

// 배열 선언시 개수는 항상 상수
