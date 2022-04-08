package main

import "fmt"

func main() {
	a := 1
	b := 1

OuterFor: // break Label
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // OuterFor 레이블까지 브레이크
			}
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

// Go에는 레이블이나 goto 문도 있으나 InstructionPointer를 강제로 바꾸기 때문에 위험하다

// 권장하지는 않음

// flag를 쓰는 것을 권장하는 이유
// 1. 코드 작성할 때 생각보다 다중 for문에서 여러단계를 break 할 일이 별로 없음
// 2. 1 같은 이유로 생각보다 코드가 별로 지저분해지지 않음
// 3. 레이블이나 goto를 지원 안하는 언어가 있기 때문에 다른 개발자들의 가독성을 위해
