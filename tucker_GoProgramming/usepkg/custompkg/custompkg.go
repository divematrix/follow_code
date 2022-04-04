package custompkg

import "fmt"

func PrintCustom() {
	fmt.Println("This is custom package!")
}

func printCustom2() {
	fmt.Println("This is custom package222!")
}

// 패키지 외부 공개

// 패키지내 타입, 전역변수, 상수, 함수, 메서드 등을 외부로 공개해서 사용할 수 있다

// 첫글자가 대문자로 시작하면 외부 공개
// 첫글자가 소문자로 시작하면 비공개
