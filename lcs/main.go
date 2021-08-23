package main

import "fmt"

func lcs(first, second string) string {
	runeOfFirst := []rune(first)
	runeOfSecond := []rune(second)
	lenFirst := len(runeOfFirst)
	lenSecond := len(runeOfSecond)

	matrix := make([][]int, lenFirst+1)
	for i := 0; i < lenFirst+1; i++ {
		matrix[i] = make([]int, lenSecond+1)
	}

	for i := 0; i < lenFirst; i++ {
		for j := 0; j < lenSecond; j++ {
			if runeOfFirst[i] == runeOfSecond[j] {
				matrix[i+1][j+1] = matrix[i][j] + 1
			} else if matrix[i+1][j] > matrix[i][j+1] {
				matrix[i+1][j+1] = matrix[i+1][j]
			} else {
				matrix[i+1][j+1] = matrix[i][j+1]
			}
		}
	}
	fmt.Println(matrix)

	readMatrix := make([]rune, 0, matrix[lenFirst][lenSecond])
	for i, j := lenFirst, lenSecond; i != 0 && j != 0; {
		if matrix[i][j] == matrix[i-1][j] {
			i--
		} else if matrix[i][j] == matrix[i][j-1] {
			j--
		} else {
			readMatrix = append(readMatrix, runeOfFirst[i-1])
			i--
			j--
		}
	}

	for i, j := 0, len(readMatrix)-1; i < j; i, j = i+1, j-1 {
		readMatrix[i], readMatrix[j] = readMatrix[j], readMatrix[i]
	}

	return string(readMatrix)
}

func main() {
	res := lcs("1234", "abcdef")
	fmt.Println(res)
}
