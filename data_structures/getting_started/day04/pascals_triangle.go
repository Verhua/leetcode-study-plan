package day04

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i, _ := range result {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}
	return result
}
