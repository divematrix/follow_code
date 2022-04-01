package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 중첩 for 문
