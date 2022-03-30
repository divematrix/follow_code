package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"

	// fmt 패키지의 Println 함수
	// 표준 출력 + 개행 + 구분자(스페이스)
	fmt.Println(msg, a)
}

// 변수 선언
// var 변수명 타입 = 초기값

// 변수 선언할 때 초기값 대입하지 않으면
// 타입에 지정된 기본값으로 변수에 초기값을 대입

// 대입연산자 : 좌항의 변수에 우항의 값을 대입

/* 변수의 4가지 속성 */
// 이름 : 개발자가 지정
// 값
// 주소 : &변수
// 타입 > 메모리크기 > 값의 범위
