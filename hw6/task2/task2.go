package main

import "fmt"

func main() {
	fmt.Println(findGCD(6, 4))
	fmt.Println(findGCD(12, 6))
}

func findGCD(a, b int) int {
	if a%b == 0 {
		return b
	}
	if b%a == 0 {
		return a
	}

	if a > b {
		return findGCD(a%b, b)
	} else {
		return findGCD(a, b%a)
	}
}
