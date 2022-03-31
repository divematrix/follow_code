package main

import "fmt"

// 타입 없는 상수 : 그냥 literal 로 컴파일되어 박힌다
const PI = 3.14

// float64 타입 상수
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100 // 오류 발생하지 않는다
	// var b int = FloatPI * 100 // 타입 오류 발생

	fmt.Println(a)
	// fmt.Println(b)
}

// 상수는 좌변으로 사용할 수 없다
// 상수는 메모리공간(동적공간)을 사용하지 않는다
// 상수는 코드에 박혀있다
// 컴퓨터 입장에서는 상수명 쓸데없다
