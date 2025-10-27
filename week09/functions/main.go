package main

/*
// 안에 숫자 바꿔보기 (애러 확인)
import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil
}

func main() {
	amount, err := paintNeeded(5.2, 3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	amount, err = paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
}
*/
/*
import (
	"fmt"
	"math" // 수학 관련 함수들이 모두 들어있음
)

func main() {
	fmt.Println(math.Sqrt(16.0))
	fmt.Println(math.Sqrt(-16.0)) // NAN : 음수 값은 구할 수 없음
}
*/
/*
import "fmt"

// 두 수 교환 함수 (포인터 X)
func swap(first int, second int) {
	temp := first
	first = second
	second = temp

	fmt.Println(first, second)
}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	swap(a, b) // pass by value : 값을 보내 원본 변경 X
	fmt.Println(a, b)
}
*/

import "fmt"

// 두 수 교환 함수 (포인터 O)
func swap(first *int, second *int) {
	temp := *first
	*first = *second
	*second = temp

	fmt.Println(*first, *second)
}

func main() {
	a, b := 10, 20
	fmt.Println(a, b)

	swap(&a, &b) // pass by pointer : 주소를 보내 원본 변경
	fmt.Println(a, b)
}