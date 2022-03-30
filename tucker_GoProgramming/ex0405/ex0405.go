package main

import "fmt"

func main() {
	var a int16 = 3456   // a는 int16타입 - 2바이트 정수
	var b int8 = int8(a) // int16 -> int8

	fmt.Println(a, b)
}

/* 타입 변환 주의 사항 */

// 큰 범위의 타입의 값을 작은 범위를 갖는 타입으로 변환할 때
// 변환되려는 타입의 범위를 넘어서는 값이 작은 범위에 다 담기지 않기 때문에 overflow bit가 발생하며 데이터 손상이 일어난다.
