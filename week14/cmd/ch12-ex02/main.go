package main

import "fmt"

func main() {
	defer fmt.Println("프로그램 종료 - 종료작업")
	fmt.Println("1. 프로그램 시작")

	panic("심각한 에러 발생!")
	// 중간에 panic이 발생해도 반드시 defer는 실행

	fmt.Println("2. 이 줄은 실행되지 않음")
}