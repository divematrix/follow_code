package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}

// 실수 오차
// 컴퓨터에서는 실수를 부동소수점으로 표현하기 때문에 정확하게 표현 불가능

// 수학의 계산과 컴퓨터의 계산은 다르다
