package main

func FindOdd(seq []int) int {
	var counter int
	for i := 0; i < len(seq); i++ {
		for j := 0; j < len(seq); j++ {
			if seq[i] == seq[j] {
				counter++
			}
		}
		if counter%2 != 0 {
			return seq[i]
		}
	}
	return 0
}
