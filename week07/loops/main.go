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
/*
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
*/
/*
import "fmt"

func main() {
	//shadowing
	// var int int = 99 // 예약어를 변수이름으로 사용하면 원래 기능이 사라짐 (shadowing)
	// var b int = 8

	// var fmt string = "inha" // 이것도 마찬가지, fmt의 기능이 가려져서 import "fmt"(fmt 정의)도 안됨
	// fmt.Println(fmt)
	fmt.Println()
}
*/
/*
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("점수 입력")

	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 있는 양쪽 여백이나 공백 제거(Trim)
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수 타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, status)
}
*/
/*
import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Seed(55) // 옛날) seed값이 같으면 같은 값이 나옴, 지금은 사용 X
	dice := rand.Intn(6) + 1 // 1 ~ 6, + 1이 없다면 범위는 0 ~ 5
	fmt.Println(dice)
}
*/

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// guess challenges players to guess a random number.
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
