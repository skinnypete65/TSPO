package main

import "fmt"

func main() {
	fmt.Println(calcFactorial(5))
}

func calcFactorial(num int) int {
	res := 1

	for num >= 2 {
		res *= num
		num--
	}

	return res
}
