package main

import "fmt"

func main() {
	// 쌍따옴표
	poet1 := "죽는 날까지 하늘을 우러러\n한 점 부끄럼이 없기를,\n잎새에 이는 바람에도\n나는 괴로워 했다.\n"

	// 홑따옴표 : 문자 하나

	// 백쿼티 : 엔터나 특수문자 그대로 출력
	poet2 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`

	fmt.Println(poet1)
	fmt.Println(poet2)
}
