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
