package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("false &&")
	}

	if true && IncreaseAndReturn() < 5 {
		fmt.Println("true &&")
	}

	if false || IncreaseAndReturn() < 5 {
		fmt.Println("false ||")
	}

	if true || IncreaseAndReturn() < 5 {
		fmt.Println("true ||")
	}

}

// Short Circuit
// && 연산자 좌항이 false면 우항을 체크할 필요없이 false값을 반환하면 되니까 우항을 실행하지 않는다
// || 연산자 좌항이 true면 우항을 체크할 필요없이 true값을 반환하면 되니까 우항을 실행하지 않는다
