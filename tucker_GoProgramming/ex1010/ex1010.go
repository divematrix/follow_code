package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a != 1, 2, 3")
	}
}

// break 키워드를 적어도 되고 안해도 된다
// Go 언어는 break 모두 적는 효과를 내준다

// 다른 언어처럼 밑 명령어도 실행하고 싶을 때는 fallthrough 키워드 쓰면 된다 (비추천)
