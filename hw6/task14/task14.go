package main

import "fmt"

func main() {
	fmt.Println(digitSqrt(9875))
}

func digitSqrt(num int) int {
	if num <= 9 {
		return num
	}

	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}

	return digitSqrt(sum)
}
