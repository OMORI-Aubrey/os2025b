package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Floor: 내림 Ceil: 올림 Round: 반올림
	fmt.Println(math.Round(2.91))
	// 단어 첫 글자가 대문자가 됨(곧 사라질 함수)
	fmt.Println(strings.Title("go developer~"))
}
