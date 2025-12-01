package main

/*
import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil { // 파일이 없으면 오픈 실패
		fmt.Println("파일 오픈 실패:", err)
		return
	}

	// defer로 파일 닫기를 미리 예약
	// defer를 쓰면 지연했다가 마지막에 실행
	defer file.Close()

	// 파일 작업 수행
	fmt.Println("파일 읽는 중...")
}
*/

import "fmt"

func main() {
	defer fmt.Println("1st defer") // 4
	defer fmt.Println("2nd defer") // 3
	defer fmt.Println("3rd defer") // 2
	fmt.Println("main logic")      // 1
}
// defer가 여러개면 LIFO 형태
// 선언한 것의 역순으로 동작
