package main

import "fmt"

func main() {
	fmt.Println(getPrimeNums(20))
}

func getPrimeNums(n int) []int {
	nums := make([]bool, n)
	primes := make([]int, 0)

	for i := 2; i < len(nums); i++ {
		if !nums[i] {
			primes = append(primes, i)
		}

		for j := i; j < len(nums); j += i {
			nums[j] = true
		}
	}

	return primes
}
