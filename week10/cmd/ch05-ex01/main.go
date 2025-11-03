package main

/*
import (
	"fmt"
)

func main() {
	// 크기가 3인 배열 선언
	var arrayBool [3]bool
	var arrayInt [3]int

	fmt.Println(arrayBool[1]) // zero value (false)

	arrayInt[1]++
	arrayInt[1] += 1 // zero value + 1

	fmt.Println(arrayInt[1])
}
*/
/*
import (
	"fmt"
)

func main() {
	// 크기가 3인 배열 선언
	// var arrayBool [3]bool = [3]bool{true, false, true} // 배열 리터럴
	arrayBool := [3]bool{true, false, true}
	var arrayInt [3]int

	fmt.Println(arrayBool[1])

	arrayInt[1] = 2
	fmt.Println(arrayInt[1])
}
*/
/*
import (
	"fmt"
	"reflect"
)

func main() {
	// 크기가 3인 배열 선언
	// var arrayBool [3]bool = [3]bool{true, false, true} // 배열 리터럴
	arrayBool := [3]bool{true, false, true}
	arrayInt := [3]int{-9, 11, 7}

	for i := 0; i < 3; i++ {
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}

	fmt.Printf("%#v\n", arrayBool)
	fmt.Printf("%#v\n", arrayInt) // 안에 원소까지 보여줌
	fmt.Println(reflect.TypeOf(arrayInt)) // 타입만 보여줌
}
*/

import (
	"fmt"
)

func main() {
	// arrayBool := [3]bool{true, false} // 값 안넣어주면 자동으로 zero value
	arrayBool := [2]bool{true, false}
		arrayInt := [3]int{-9, 11, 7}

	// for i := 0; i < 3; i++ { // 없는 인덱스의 값을 찍으면 에러
	// for i := 0; i < len(arrayInt); i++ { // 둘이 길이가 달라서 에러
	// for i := 3; i < len(arrayBool); i++ { // for 문이 돌지도 않음
	// for i := 0; i < len(arrayBool); i-- { // 증감이 잘못되어 프로그램 터짐
	for i := 0; i < len(arrayBool); i++ { 
		fmt.Println(i, arrayBool[i])
		fmt.Println(i, arrayInt[i])
	}

}
