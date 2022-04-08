package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for {
		time.Sleep(time.Second)
		fmt.Println(i)
		i++
	}
}

// 무한루프 - for문에서 조건문이 모두 true 일때 : 조건문을 생략하면 true로 자동 설정
