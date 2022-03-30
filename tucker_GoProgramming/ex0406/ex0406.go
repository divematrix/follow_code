package main

import "fmt"

// 패키지 전역 변수 g
// 패키지 전역에 살아있는 변수
var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// m = s + 20
	// 변수 s의 범위를 벗어낫기 때문에 에러
	// 변수 s가 선언된 위치의 중괄호 안에서만 s가 유지
}
