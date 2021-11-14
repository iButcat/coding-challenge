package kata

import "fmt"

func sortWord(ele string) string {
	var res = []byte(ele)
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[i] < res[j] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	return string(res)
}

func Anagrams(word string, words []string) []string {
	var res []string
	for i := 0; i < len(words); i++ {
		if sortWord(word) == sortWord(words[i]) {
			res = append(res, words[i])
		}
	}
	fmt.Println("res: ", res)
	return res
}
