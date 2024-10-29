package main

import "fmt"

func main() {
	arr := []int{1, 5, 8, 10, 11, 2, 0}
	fmt.Println(findMinMax(arr))
}

func findMinMax(arr []int) (int, int) {
	if len(arr) == 0 {
		return 0, 0
	}
	minNum, maxNum := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		minNum = min(minNum, arr[i])
		maxNum = max(maxNum, arr[i])
	}

	return minNum, maxNum
}
