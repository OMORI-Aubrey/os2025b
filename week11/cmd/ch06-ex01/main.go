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
/*
import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[1:3] // slicing

	for _, subject := range subjects {
		fmt.Println(subject)
	}

	fmt.Println("=============")

	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
*/
/*
import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3]
	
	// subjectsSlice[0] = "Java" // 원본 값 변경
	subjects[0] = "Java" // 메모리 주소 공유 (얕은 복사)

	for _, subject := range subjects {
		fmt.Println(subject)
	}

	fmt.Println("=============")

	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
*/

import "fmt"

func main() {
	subjects := [4]string{"Go", "Javascript", "Python", "Linux"}
	subjectsSlice := subjects[:3]

	subjects[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go") // 배열의 길이가 4기 때문에 Linux가 덮어씌어짐

	for _, subject := range subjects {
		fmt.Println(subject)
	}

	fmt.Println("=============")

	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}