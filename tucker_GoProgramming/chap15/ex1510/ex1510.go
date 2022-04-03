package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}

// 문자열 비교

// 사전식 비교 - 대문자가 더 작다

// 문자열 변수는 일종의 포인터
// 속성값1은 문자열 데이터
// 속성값2는 문자열 길이
