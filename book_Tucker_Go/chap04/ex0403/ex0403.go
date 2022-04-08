package main

import "fmt"

func main() {
	// 변수 선언 방법1 - 기본
	var a int = 3

	// 변수 선언 방법2 - 타입에 지정된 기본값으로 초기값 자동 대입
	var b int

	// 변수 선언 방법3 - 초기값에 맞는 타입으로 자동 설정
	var c = 4

	// 변수 선언 방법4 - 선언대입문(var, 타입 생략)
	d := 5

	var e = "Hello"
	f := 3.14

	fmt.Println(a, b, c, d, e, f)
}
