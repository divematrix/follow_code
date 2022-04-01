package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
	case 2:
		fmt.Println("a == 2")
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fmt.Println("a == 3")
		fmt.Println("a == 3")
	default:
		fmt.Println("a != 1, 2, 3", a)
	}
}

// if 문 : 조건을 검사
// switch 문 : 값을 검사

// switch 초기문; 비교값 {
//  case 값1:
//   코드
//  default:
//   코드
// }
