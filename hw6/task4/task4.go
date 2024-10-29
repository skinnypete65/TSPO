package main

import "fmt"

func main() {
	multTable := calcMultTable()
	printMat(multTable)
}

func calcMultTable() [][]int {
	ans := make([][]int, 10)

	for i := 0; i < 10; i++ {
		ans[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			num1, num2 := i+1, j+1

			ans[i][j] = num1 * num2
		}
	}

	return ans
}

func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			fmt.Printf("%d\t", mat[i][j])
		}
		fmt.Println()
	}
}
