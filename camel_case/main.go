package main

import (
	"fmt"
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	var res string
	matchToRemove := regexp.MustCompile(`(-|_)`)
	res = string(matchToRemove.ReplaceAll([]byte(s), []byte(" ")))
	if strings.ToUpper(string(res[0])) == string(res[0]) {
		res = strings.Title(res)
		res = strings.ReplaceAll(res, " ", "")
	} else {
		res = strings.Title(res)
		res = strings.ReplaceAll(res, " ", "")
		res = strings.ToLower(string(res[0])) + string(res[1:])
	}
	return res
}

func main() {
	fmt.Println("Hello")
}

//str = strings.ToLower(string(str[0])) + string(str[1:])
