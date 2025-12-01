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
