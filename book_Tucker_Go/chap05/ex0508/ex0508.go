package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	// 첫 입력 : 숫자를 입력 받기로 약속되어 졌지만 일부러 에러를 발생 시켜야 하므로 문자열을 넣어주자
	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	// 두번재 입력 : 첫 입력이 에러 발생후 버퍼가 읽어와 처리됬음을 확인하기 위해 숫자 두개를 입력해주자
	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}

/* 표준 입력 버퍼의 남은 문자열 읽어오기 */

// 소스는 int 형식의 데이터 2개를 입력 받는 코드
// 문자열이 입력되면 숫자가 아니므로 에러가 발생
// 입력 받은 첫 문자가 에러가 나면 뒤의 입력은 입력 버퍼에서 처리되지 않고 남아있게 된다

// 입력 버퍼의 남은 데이터를 읽어와서 처리할 필요가 있다
// 입력 버퍼에서 개행문자가 나올때까지 읽어오는 코드를 짜서
// 입력 버퍼가 새로운 입력을 받을 수 있도록 비워주자
