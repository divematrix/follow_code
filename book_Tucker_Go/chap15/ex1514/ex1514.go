package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1))
	// 1. unsafe 패키지의 Pointer타입으로 변환
	// 2. reflect.StringHeader 타입으로 변환
	// string의 데이터 메모리 주소값을 반환

	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)
	// str1 과 str2의 값의 메모리 주소값이 동일 확인
}

// string은 immutable : 문자열 일부만 수정을 할 수없다
