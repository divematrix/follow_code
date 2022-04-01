package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("입력하세요")

		var number int

		// _ 빈칸 지시자
		// Go에서는 변수를 선언하면 반드시 써야하는데
		// 함수의 반환값중 영원히 안쓸 값에 빈칸 지시자를 써서 안쓴다는 표시를 할 수 있다.
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자로 입력하세요")

			// 키보드 버프를 지웁니다
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d입니다.\n", number)

		if number%2 == 0 {
			break // 짝수 검사를 합니다.
		}

	}

	fmt.Println("for 문이 종료되었습니다.")
}
