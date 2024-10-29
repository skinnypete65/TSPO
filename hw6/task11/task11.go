package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isArmstrongNum(153))
	fmt.Println(isArmstrongNum(154))
}

func isArmstrongNum(num int) bool {
	temp := num

	size := 0

	for temp > 0 {
		temp /= 10
		size++
	}

	digitSum := 0
	temp = num

	for temp > 0 {
		digit := temp % 10
		temp /= 10

		digitSum += int(math.Pow(float64(digit), float64(size)))
	}

	return num == digitSum
}
