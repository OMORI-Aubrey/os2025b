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
/*
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
*/

import (
	"bufio" // bufio: BUFfer Input Output
	"fmt"
	"log"
	"os" // os: Operating System
)

func main() {
	r := bufio.NewReader(os.Stdin)
	// return값 두 개 필요 -> 다중 리턴 가능
	i, err := r.ReadString('\n') // '' 안에 있는 글자 까지 입력받음 -> '\n' : \n까지 입력받음
	// i, _ := r.ReadString('\n') // 변수 _ 는 값을 받지만 사용하지 않겠다는 뜻(리턴값 무시)

	if err != nil { // nil: zero value (값 없음, 아무것도 가리키지 않음)
		log.Fatal(err) //에러 메시지 보고, 프로그램 종료
	}

	fmt.Println(i)
}
