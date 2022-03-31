package main

import "fmt"

func main() {
	const pi1 float64 = 3.141592653589793228 // 상수
	var pi2 float64 = 3.141592653589793228   // 변수

	// pi1 = 3 : 상수에는 새로운 대입 못한다
	pi2 = 4

	fmt.Printf("원주율: %f\n", pi1)
	fmt.Printf("원주율: %f\n", pi2)
}

// 상수
// 상수는 한번 선언하면 값을 바꿀 수 없다

// 상수는 변하지 않는 값에 사용
// 상수는 코드값으로 사용 ex) ASCII 코드
