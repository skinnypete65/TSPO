package main

import "fmt"

func main() {
	fmt.Println(sumOfNumsSlice([]int{1, 2, 3, 4, 5}))
}

func sumOfNumsSlice(sl []int) int {
	sum := 0
	for _, num := range sl {
		sum += num
	}

	return sum
}
