package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

// 구조체를 포함하는 구조체
type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

// 포함된 필드 방식 embedded field
// (상속 받는 방식)
type VIPUser2 struct {
	User
	VIPLevel2 int
	Price2    int
	Name      string // 포함된 필드 방식으로 구조체를 포함할 경우 속성에 포함된 구조체의 속성명과 같은 속성명을 가질때
	// 새로 정의한 속성명이 우선권
	// 포함된 구조체의 속성을 호출하려면 변수명.포함된구조체명.속성명으로 호출
}

func main() {
	user := User{"송하나", "hana", 23}

	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
	}

	vip2 := VIPUser2{
		User{"순신", "sunsin", 50},
		2,
		100,
		"영미",
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)

	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n", vip.UserInfo.Name, vip.UserInfo.ID, vip.UserInfo.Age, vip.VIPLevel, vip.Price)

	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP레벨: %d VIP가격: %d만원\n", vip2.Name, vip2.ID, vip2.Age, vip2.VIPLevel2, vip2.Price2)
}

// 구조체의 사이즈는 모든 필드의 사이즈를 더한 크기
