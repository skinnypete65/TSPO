package main

import "fmt"

func main() {
	fmt.Println(getMaxNum(4, 9, 7))
}

func getMaxNum(num1, num2, num3 int) int {
	max1 := max(num1, num2)
	return max(max1, num3)
}
