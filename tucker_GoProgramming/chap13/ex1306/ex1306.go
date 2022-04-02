package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
} // 4+8 = 12바이트

func main() {
	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user))
}

// unsafe.Sizeof(데이터) 데이터의 메모리 사이즈

// 왜 16바이트가 나올까?
// 메모리 정렬(Memory Alignment)

// 64비트 컴퓨터는 레지스터 크기가 8비트
// 4바이트 크기 타입 이후
// 8바이트 크기 타입이 오면
// 레지스터 읽어오기 편하게 4바이트를 비우고
// 레지스터에 담기 편하게 다음 타입을 저장
// : Memory padding

// 구조체가 8 + 4 바이트 크기면 12로 나옴
