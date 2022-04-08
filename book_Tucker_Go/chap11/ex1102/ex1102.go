package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println("6 *", i, "=", 6*i)
	}
}

// continue 키워드 : 후처리문으로 건너뜀
// break 키워드 : 반복문을 종료
