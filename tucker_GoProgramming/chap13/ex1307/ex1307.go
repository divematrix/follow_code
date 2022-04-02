package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	A int8 // 1
	C int8 // 1
	E int8 // 1
	B int  // 8
	D int  // 8
}

// 1 + 8 + 1 + 8 + 1 = 19 -> 메모리 패딩으로 40
// 1 + 1 + 1 + 8 + 8 = 19 -> 메모리 패딩으로 24
// 메모리가 작은 환경에서는 작은 타입을 몰아서 작성하자

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}

// 구조체의 역할 - 프로그램의 방향
// 결합도(의존성)은 낮게 응집도는 높게
// low coupling, high cohesion

// - 함수는 관련 코드 블록을 묶어서 응집도를 높이고 재사용성을 증가
// - 배열은 같은 타입의 데이터들을 묶어서 응집도를 높임
// - 구조체는 관련 데이터들을 묶어서 응집도를 높이고 재사용성을 증가

// 구조체는 객체 지향 프로그래밍의 기반이 됨
