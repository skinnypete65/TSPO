package main

import "fmt"

func main() {
	fmt.Println(getFibonacciNums(7))
}

func getFibonacciNums(n int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	
	num1, num2 := 1, 0
	ans := make([]int, 0)
	ans = append(ans, num2, num1)

	for i := 2; i < n; i++ {
		sum := num1 + num2
		ans = append(ans, sum)

		num1, num2 = sum, num1
	}

	return ans
}
