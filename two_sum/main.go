package main

func TwoSum(numbers []int, target int) [2]int {
	var tupleRes = [2]int{}
	var length int = len(numbers)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if (numbers[i] + numbers[j]) == target {
				tupleRes[0], tupleRes[1] = i, j
				return tupleRes
			}
		}
	}
	return tupleRes
}
