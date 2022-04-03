package main

import "fmt"

func main() {
	str := "Hello 월드"
	arr := []rune(str)
	// rune : int32의 별칭 타입
	// []rune으로 타입 변환후 한글자씩 순회

	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자값:%c\n", arr[i], arr[i], arr[i])
	}

}

// len(배열) : 배열의 개수 반환
