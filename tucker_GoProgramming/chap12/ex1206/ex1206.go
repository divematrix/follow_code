package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, // 중괄호가 다른줄에서 닫히면 , 컴마 찍어줘야 함
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println("")
	}
}

// 다중배열
// [3차][2차][1차]int 식으로 정의

// Go에서 배열 크기 = 타입크기 * 항목 개수
