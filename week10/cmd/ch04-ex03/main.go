package main

import (
	"fmt"
	"log"
	"week10/pkg/keyboard"
)

func main() {
	fmt.Print("화씨 온도 입력: ")

	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Printf("화씨 %0.2f는 섭씨 %0.2f도 입니다.\n", fahrenheit, celsius)

}
