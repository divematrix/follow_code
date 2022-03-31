package main

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	// 문자열 형식 - 5칸안에 출력
	fmt.Printf("%5d, %5d\n", a, b)

	// 문자열 형식 - 5칸안에 출력, 빈칸에 0채우기
	fmt.Printf("%05d, %05d\n", a, b)

	// 문자열 형식 - 5칸안에 출력, 왼쪽 정렬
	// 빈칸에 0을 채워 값이 변화하면 0을 채우지 않는다
	fmt.Printf("%-5d, %-05d\n", a, b)

	// 문자열 형식 - 5칸안에 출력
	// 값이 지정된 칸을 넘게 되면 칸 지정은 취소
	fmt.Printf("%5d, %5d\n", c, c)
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-5d, %-05d\n", c, c)
}

// 실수 포맷팅
// %0.15f - 실수를 15자리까지 표현
