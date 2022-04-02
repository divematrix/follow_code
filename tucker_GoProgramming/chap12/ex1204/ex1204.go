package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	for i, v := range t {
		fmt.Println(i, v)
	}
}

// range 키워드
// range 자료구조 : 각 요소들을 순회 - 값을 두개 반환
// ->  배열 : 인덱스, 요소값
