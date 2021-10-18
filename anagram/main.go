package main

import (
	"fmt"
	"strings"
)

func sortStrings(str string) string {
	lower := strings.ToLower(str)
	var convertStr = []byte(lower)
	for i := 0; i < len(convertStr); i++ {
		for j := i + 1; j < len(convertStr); j++ {
			if convertStr[i] < convertStr[j] {
				convertStr[i], convertStr[j] = convertStr[j], convertStr[i]
			}
		}
	}
	return string(convertStr)
}

func IsAnagram(str1, str2 string) bool {
	return sortStrings(str1) == sortStrings(str2)
}

func main() {
	fmt.Println(IsAnagram("hello", "olleh"))         // true
	fmt.Println(IsAnagram("eas", "eax"))             // false
	fmt.Println(IsAnagram("excalibur", "rubilacxe")) // true
}
