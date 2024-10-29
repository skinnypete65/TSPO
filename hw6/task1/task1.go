package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPrime(29))
	fmt.Println(isPrime(12))
	fmt.Println(isPrime(11))
	fmt.Println(isPrime(16))
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	sqrtNum := int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrtNum; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}
