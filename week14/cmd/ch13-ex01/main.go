package main

/*
import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(200 * time.Millisecond)
	}
}


func main() {
	// say("고루틴 키워드 뺌") // 하나 먼저 다 하고 다음꺼 실행
	go say("고루틴") // 새 고루틴에서 실행
	say("메인")     // 메인 고루틴에서 실행
}
*/
/*
import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(200 * time.Millisecond)
	}
}


func main() {
	start := time.Now()

	// say("고루틴 뺌")
	go say("고루틴") // 비동기 처리
	say("메인")

	fmt.Println("전채 실행 시간:", time.Since(start))
}
*/

import (
	"fmt"
	"time"
)

func say(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(200 * time.Millisecond)
	}
}


func hi(msg string) {
	time.Sleep(8000 * time.Millisecond)
	fmt.Println(msg)
}


func main() {
	start := time.Now()

	// go 키워드를 둘다 줘버리면 아무것도 안나옴
	// 프로그램 끝날때까지 기다려주지 않기 때문
	go say("고루틴1") // 6초 걸림 (이 일이 먼저 끝남)
	go hi("고루틴2") // 8초 걸림

	time.Sleep(10 * time.Second) // 메인 고루틴 대기 시간 줘서 해결
	fmt.Println("전채 실행 시간:", time.Since(start))
}