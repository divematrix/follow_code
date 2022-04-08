package main

import (
	htemplate "html/template" // 별칭 "패키지명" - 별칭으로 작동
	"text/template"           // 같은 이름의 패키지 - 밑에있는 패키지가 작동
)

func main() {
	template.New("foo").Parse(`{{define "T"}}Hello`)
	htemplate.New("foo").Parse(`{{define "T"}}Hello`)
}

// 모듈은 패키지의 모음
// 모듈은 1개 이상 패키지를 포함

// 패키지란 코드를 묶는 단위

// 프로그램이란
// 실행 시작 지점을 포함한 패키지
// 즉 main 함수를 포함한 main 패키지

// 패키지는 소스에서 / 이후 마지막 이름으로 작동
// 동일 이름의 패키지를 import 하면 마지막에 적어준 패키지가 작동

// 패키지를 import 할 때 별칭을 붙여줄 수 있다
// 별칭 "패키지명"

// import 한 패키지는 써야 컴파일에서 에러가 발생하지 않는다.
// 하지만 부수효과 때문에 부득이 import 할때는 별칭으로 빈칸지시자 _ 를 쓴다
// 패키지 초기화에 따른 부가 효과를 위해서.
