package main

/*
import "fmt"

func main() {
	ch := make(chan int) // int 채널 생성

	go func() {
		ch <- 123 // 채널에 값 보내기
	}()

	x := <-ch // 채널에서 값 받기
	fmt.Println(x)
}
*/
/*
import "fmt"

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go abc(channel1)
	go def(channel2)
	
	// 순서가 채널별로 동기화
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Print(<-channel1)
	fmt.Print(<-channel2)
	fmt.Println()
}
*/

import "fmt"

func abc(channel chan string) {
	channel <- "내고향 스페셜\n"
	channel <- "KBS 뉴스광장\n"
	channel <- "인간극장\n"
}

func def(channel chan string) {
	channel <- "건강의 재구성 설록(재)\n"
	channel <- "오늘N\n"
	channel <- "찾아가는 꾸러기 교실\n"
}

func main() {
	kbs := make(chan string)
	mbc := make(chan string)
	
	go abc(kbs)
	go def(mbc)

	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Print(<-kbs)
	fmt.Print(<-mbc)
	fmt.Println()
}