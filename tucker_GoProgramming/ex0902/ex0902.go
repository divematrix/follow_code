package main

import "fmt"

func main() {
	temp := 33

	if temp > 28 {
		fmt.Println("에어컨을 켠다")
	} else if temp <= 3 {
		fmt.Println("히터를 켠다")
	} else if temp <= 18 {
		fmt.Println("나가자!")
	} else {
		fmt.Println("덥다")
	}
}

// 조건문 if문
// if 조건 { 코드블럭 }      : 필수
// else if 조건 { 코드블럭 } : 옵션, 중  간, 여러번
// else { 코드블럭 }         : 옵션, 마지막, 한번만
