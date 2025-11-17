package main

/*
import (
	"fmt"
	"log"
	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	var names []string
	var counts []int

	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		// 처음 나온 이름
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Println(name, ":", counts[i])
	}
}
*/

import (
	"fmt"
	"log"
	"github.com/headfirstgo/datafile"
)

// map을 사용하여 개선한 버전
func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}

	for name, count := range counts {
		fmt.Println(name, ":", count)
	}
}