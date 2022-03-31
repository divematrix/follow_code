package main

import "fmt"

func main() {
	var x int8 = 16
	var y int8 = -128
	var z int8 = -1
	var w uint8 = 128

	fmt.Printf("x: %08b x>>2: %08b x>>2 : %d\n", x, x>>2, x>>2)
	fmt.Printf("y: %08b y>>2: %08b y>>2 : %d\n", uint8(y), uint8(y>>2), y>>2)
	fmt.Printf("z: %08b z>>2: %08b z>>2 : %d\n", uint8(z), uint8(z>>2), z>>2)
	fmt.Printf("w: %08b w>>2: %08b w>>2 : %d\n", w, w>>2, w>>2)
}

// 오른쪽 시프트 연산자
// 대상 숫자에 2의 승수만큼 나누는 역할을 하지만

// 양수와 음수에서 앞에 채워지는 숫자가 0, 1로 다르기 때문에 주의해야 한다
