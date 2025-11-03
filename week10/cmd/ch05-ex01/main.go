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