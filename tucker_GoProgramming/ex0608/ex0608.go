package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))
}

// math.Nextafter() : 수학계산에서 오차 보정
// 두 피연산자가 1비트이내 차이나면 true 반환
