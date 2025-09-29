package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
		// Floor: 내림, Ceil: 올림, Round: 반올림
		fmt.Println(math.Round(2.91))
		// 단어 첫 글자가 대문자가 됨(곧 사라질 함수)
		fmt.Println(strings.Title("go developer~"))
		fmt.Println("Kim\nInha\t\"\\") // C언어랑 비슷함
		fmt.Println('A', '가')          // Rune (유니코드)
	*/
	/*
		fmt.Println(reflect.TypeOf(2.31)) // float64
		fmt.Println(reflect.TypeOf("go developer~")) //string
		fmt.Println(reflect.TypeOf('A')) //int32
		fmt.Println(reflect.TypeOf(true)) //bool
		fmt.Println(reflect.TypeOf(91)) //int
	*/
	/*
		// var name string
		// var id int
		// name = "Kim Inha"
		// id = 1000
		// var name string = "Kim Inha"
		// var id int = 1000

		// 타입 자동 추론으로 간편하게 선언
		name := "Kim Inha"
		id := 1000

		// 선언은 두 번 이상 할 수 없음
		// id := 1000 () // 이미 위에 선언
		id = 1000

		// gpa := 3.99 // 자동으로 float64 (double)
		var gpa float32 = 3.99 // 내가 명시한 자료형으로 선언

		fmt.Println(name, reflect.TypeOf(name))
		fmt.Println(id, reflect.TypeOf(id))
		fmt.Println(gpa, reflect.TypeOf(gpa))
	*/

	// Zero values
	var f64 float64
	var i16 int16
	var t bool
	var s string
	var i int

	fmt.Println(f64, reflect.TypeOf(f64))
	fmt.Println(i16, reflect.TypeOf(i16))
	fmt.Println(t, reflect.TypeOf(t))
	fmt.Println(s, reflect.TypeOf(s))
	fmt.Println(i, reflect.TypeOf(i))
	
}
