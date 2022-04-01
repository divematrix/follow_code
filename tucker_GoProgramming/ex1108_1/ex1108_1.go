package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true
				break // 두번째 for문만 빠져나감
			}
		}
		if found {
			break // 두번재 for문 break 빠져나갔을때 첫번째 반복문 까지 빠져나가기 위해 쓰는 flag 변수
		}
	}

	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
