package main

import (
	"fmt"
	"sort"
)

func HighestRank(nums []int) int {
	var mostRepeated = make(map[int]int, len(nums))
	var max int
	var res int

	sort.Ints(nums)
	for _, v := range nums {
		mostRepeated[v]++
		if mostRepeated[v]+1 > max {
			max = mostRepeated[v]
			res = v
		}
	}
	return res
}

func main() {
	fmt.Println(HighestRank([]int{12, 10, 8, 12, 7, 6, 4, 10, 12}))     // 12
	fmt.Println(HighestRank([]int{10, 12, 8, 12, 7, 6, 4, 10, 12, 10})) // 12
	fmt.Println(HighestRank([]int{10, 10, 20, 20}))                     // 20
}
