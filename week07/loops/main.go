package main

/*
import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 3.2
	var width int = 2

	// 타입이 다르면 연산이 안됨 -> 타입 통일 시키기
	// fmt.Println("면적은", length * width)
	fmt.Println("면적은", int(length) * width)
	fmt.Println("length > width?", int(length) > width)
	fmt.Println("형변환", reflect.TypeOf(int(length)))
	fmt.Println("원본", reflect.TypeOf(length))
}
*/
/*
import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	// var year int = now.Year

	// month := now.Month()
	var month time.Month = now.Month() // 타입을 string으로 하면 타입이 맞지 않아 오류

	var day int = now.Day()

	fmt.Println(month, day)
}
*/

import (
	"fmt"
	"strings"
)

func main() {
	univ := "Go$ inha$"
	changer := strings.NewReplacer("$", "!")
	changed := changer.Replace(univ)

	fmt.Println(changed)
}
