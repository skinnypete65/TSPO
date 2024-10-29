package main

import (
	"fmt"
	"math/rand"
)

// # - for dead
// . - for alive

const (
	rows = 10
	cols = 10
)

func main() {
	startGame(2)
}

func startGame(steps int) {
	grid := initGrid()
	printGrid(grid)

	for i := 0; i < steps; i++ {
		nextState(grid)
		printGrid(grid)
	}
}

func initGrid() [][]bool {
	grid := make([][]bool, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]bool, cols)

		for col := 0; col < cols; col++ {
			grid[row][col] = rand.Intn(2) == 0
		}
	}

	return grid
}

func nextState(grid [][]bool) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			neigsCnt := countAliveNeigs(grid, col, row)

			if grid[row][col] {
				grid[row][col] = neigsCnt == 2 || neigsCnt == 3
			} else {
				grid[row][col] = neigsCnt == 3
			}
		}
	}
}

func countAliveNeigs(grid [][]bool, row, col int) int {
	// first - for row, second - for col
	// (-1, -1) (-1, 0) (-1, 1)
	// (0, -1)  (r, c)  (0, 1)
	// (1, -1)  (1, 0)  (1, 1)

	cnt := 0

	diffs := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, diff := range diffs {
		dRow, dCol := row+diff[0], col+diff[1]

		if dRow >= 0 && dRow < rows && dCol >= 0 && dCol < cols && grid[dRow][dCol] {
			cnt++
		}
	}

	return cnt
}

func printGrid(grid [][]bool) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] {
				fmt.Print(". ")
			} else {
				fmt.Print("# ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
}
