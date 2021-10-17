package main

import "fmt"

func FindUniq(arr []float32) float32 {
	var res float32
	var findUnique = make(map[float32]int)
	for i := 0; i < len(arr); i++ {
		findUnique[arr[i]]++
	}
	for i, v := range findUnique {
		if v == 1 {
			res = i
		}
	}
	return res
}

func main() {
	var arr = []float32{1, 1, 1, 2, 2, 3, 5, 5, 6, 6}
	fmt.Println(FindUniq(arr)) // 3
}
