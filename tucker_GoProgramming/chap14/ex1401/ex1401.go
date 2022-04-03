// 포인터
// 메모리 주소를 값으로 갖는 타입
// 메모리 공간을 가리키고 있다

package main

import "fmt"

func main() {
	var a int = 500
	var p *int
	p = &a

	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100 // a = 100
	fmt.Printf("a의 값: %d\n", a)
}
