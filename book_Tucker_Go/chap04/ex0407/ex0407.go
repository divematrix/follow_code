package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b // 4,266,663.334329
	var d float32 = c * 3 // 12,799,990.002987

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// 부동소수점 실수 연산에서 오차가 발생한다.
// 고정소수점 실수 연산은 정확성을 위해 있으나 속도가 느리다
