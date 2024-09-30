package main

import "fmt"

func main() {
	fmt.Println(doubleSlice([]int{1, 2, 3, 4}))
}

func doubleSlice(src []int) []int {
	newSlice := make([]int, len(src))

	for i := range src {
		newSlice[i] = src[i] * 2
	}

	return newSlice
}
