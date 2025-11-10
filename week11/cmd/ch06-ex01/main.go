package main

/*
import "fmt"

func main() {
	// var subjects []string
	// subjects = make([]string, 3)
	// subjects := make([]string, 3)

	// subjects[0] = "Go"
	// subjects[2] = "Python"

	subjects := []string{"Go", "", "Python"} // slice literal

	for _, subject := range subjects {
		fmt.Println(subject)
	}
}
*/

import "fmt"

func main() {
	subjects := []string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[1:3] // slicing

	for _, subject := range subjects {
		fmt.Println(subject)
	}

	fmt.Println("=============")

	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}