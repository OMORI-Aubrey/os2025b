package main

import (
	"greeting"
)

func main() {
	// 우리가 직접 만든 커스텀 패키지
	// 함수 이름이 소문자로 시작하면 외부 패키지 접근 불가능
	greeting.Hello()
	greeting.Hi()
}
