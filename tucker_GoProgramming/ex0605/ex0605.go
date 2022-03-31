package main

import "fmt"

func main() {
	var x int8 = 127
	fmt.Printf("%d < %d+1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t=%4d, %08b\n", x, x)
	fmt.Printf("x+1\t=%4d, %08b\n", x+1, uint8(x+1))
}

// 정수 오버플로우
// 타입 범위에서 가장 큰 양의 정수에 +1 을 하면
// 타입 범위에서 가장 작은 음의 정수로 결과가 나온다

// 수학에서 정의와 다른 결과가 도출
