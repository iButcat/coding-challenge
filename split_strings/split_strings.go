package main

import "regexp"

func Solution(str string) []string {
	var res []string
	re := regexp.MustCompile(`..?`)
	chunks := re.FindAll([]byte(str), -1)
	for i := 0; i < len(chunks); i++ {
		if len(chunks[i]) == 1 {
			res = append(res, string(chunks[i])+"_")
		} else {
			res = append(res, string(chunks[i]))
		}
	}
	return res
}
