package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	var countDup = make(map[rune]int)
	for _, v := range strings.ToLower(s1) {
		countDup[v]++
	}
	var counter int
	for i, _ := range countDup {
		if countDup[i] > 1 {
			counter = counter + 1
		}
	}
	return counter
}

func main() {
	res := duplicate_count("abcdeaB11") // 3
	fmt.Println("Should be 3: ", res)
	res1 := duplicate_count("indivisibility") // 1
	fmt.Println("Should be 1: ", res1)
	res2 := duplicate_count("abcdeaB11")
	fmt.Println("Should be 3: ", res2)
}
