package main

import (
	"strings"
)

func High(s string) string {
	var lowercase = "abcdefghijklmnopqrstuvwxyz"
	var letterValue = make(map[rune]int, len(lowercase))
	var res string

	for index, value := range lowercase {
		letterValue[value] = index + 1
	}

	var counter int
	var mySplit = strings.Split(s, " ")
	var wordValues []int
	for i := 0; i < len(mySplit); i++ {
		for j := 0; j < len(mySplit[i]); j++ {
			counter = counter + letterValue[rune(mySplit[i][j])]
		}
		wordValues = append(wordValues, counter)
		counter = 0
	}

	var max int
	for index, value := range wordValues {
		if value > max {
			max = value
			res = mySplit[index]
		}
	}
	return res
}
