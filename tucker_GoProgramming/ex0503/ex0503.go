package main

import "fmt"

func main() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}
}

// 표준 입력 함수에는 입력받은 값을 저장할 메모리주소가 필요하다
// &변수 : Go 언어에서 변수의 메모리주소값을 나타내는 방식
// 입력 인자에 메모리주소 값을 입력 : &변수
