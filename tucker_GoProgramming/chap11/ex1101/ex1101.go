package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}

	// fmt.Println(i) - i는 for문의 지역변수 : 에러발생
}

// for 문
// for 초기문; 조건문; 후처리 { 코드 }

// 초기문, 조건문, 후처리는 생략 가능

// 조건문이 생략되면 true 조건으로 설정 - 무한루프
