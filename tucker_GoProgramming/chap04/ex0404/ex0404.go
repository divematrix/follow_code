package main

import "fmt"

func main() {
	a := 3 // int + 64bit 컴퓨터 = int64
	var b float64 = 3.5

	var c int = int(b) // 3.5 -> 3
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}

/* Go는 슈퍼 강타입 언어 */
// 연산의 각 항목의 타입은 '반드시' 같아야 한다.

// 형변환 연산
// 타입(데이터)

// 별칭 타입 생성
// type 별칭타입명 대상타입명
