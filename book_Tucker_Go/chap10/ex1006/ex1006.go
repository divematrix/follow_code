package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("age is", age) - Error : age 변수는 switch 문의 지역변수로 사라진 상태
}

// switch 초기문; 조건문 { }
