package main

import (
	"fmt"
	"tucker_GoProgramming/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/* 따라하기 Go 모듈 만들고 외부 패키지 활용하기

1. goproject/usepkg 폴더 생성
2. go mod init goproject/usepkg
3. mkdir custompkg
4. Custompkg.go
5. mkdir program
6. usepkg.go
7. go mod tidy
*/

// cmd >> go tidy
// 모듈에서 import 한 외부 패키지 다운
// gopath > pkg > sumdb  경로에
// go.sum 파일에 외부 패키지 정보를 저장한다.
