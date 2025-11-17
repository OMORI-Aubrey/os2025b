package main

import (
	"fmt"
)

type subscriber struct {
	name string
	price int
}

func applyPrice(s *subscriber)  {
	s.price = 10000
	s.name = "Park Inha"
}

func main() {
	var s1 subscriber
	// s1.name = "Kim Inha"
	applyPrice(&s1) // 포인터로 안받으면 기본값

	fmt.Println(s1.name, s1.price)
}