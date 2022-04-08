package main

import (
	"tucker_GoProgramming/ex1605/exinit"
)

func main() {
	exinit.PrintD()
}

// 패키지 초기화

// 패키지가 프로그램에 포함되어 초기화 될때
// 패키지내 init()가 한번만 호출된다.

// init() 를 통해서 패키지내 전역 변수를 초기화 한다.
