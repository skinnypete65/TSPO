package main

import "fmt"

func main() {
	triangle := generatePascalTriangle(10)
	printMat(triangle)
}

func generatePascalTriangle(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		ans[i] = make([]int, i+1)
	}

	ans[0][0] = 1

	for i := 1; i < numRows; i++ {
		prevRow := ans[i-1]
		row := ans[i]

		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = prevRow[j-1] + prevRow[j]
		}
	}

	return ans
}

func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Println()
	}
}
