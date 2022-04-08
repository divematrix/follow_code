package main

import "fmt"

type Data struct {
	value int
	data  [200]int
}

// func ChangeData(arg Data)
func ChangeData(arg *Data) {
	arg.value = 999
	// (*arg).value 와 같다
	arg.data[100] = 999
}

func main() {
	var data Data

	// ChangeData(data)
	ChangeData(&data)

	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])
}

// 포인터를 쓰는 이유
// 함수의 인자로 쓰이면 r-value(값)으로 쓰인다
// 인자로 쓰인 변수의 값을 변화 시키지 못한다.

// 인자로 쓰인 변수의 값을 함수로 변화 시키기 위해 포인터 쓴다.

// new() 내장 함수
// new(Data) : &Data{} : Data타입의 객체 생성후 메모리 주소 반환
// 차이점 : new는 객체의 속성이 모두 기본값
