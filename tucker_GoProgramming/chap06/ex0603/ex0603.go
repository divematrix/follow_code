package main

import (
	"fmt"
)

func main() {
	var x int8 = 4
	var y int8 = 64

	fmt.Printf("x:%08b x<<2:%08b x<<2: %d\n", x, x<<2, x<<2)
	fmt.Printf("y:%08b y<<2:%08b y<<2: %d\n", y, y<<2, y<<2)
}

// 왼쪽 시프트 연산자
// 대상 숫자에 2의 승수만큼 곱하는 역할을 하지만

// 범위를 벗어나는 비트는 탈락되기 때문에
// 무조건 2의 승수만큼 곱하는 결과를 나타내는것은 아니다
