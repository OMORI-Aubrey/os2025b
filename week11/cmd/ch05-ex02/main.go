// average calculates the average of several numbers.
package main

import (
	"fmt"
	"github.com/headfirstgo/datafile"
	"log"
)

func main() {
	weights, err := datafile.GetFloats("meatWeight.txt")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0
	for _, number := range weights {
		sum += number
	}
	
	weeks := float64(len(weights))
	fmt.Printf("평균 고기 소비량: %0.2f\n", sum/weeks)
}
