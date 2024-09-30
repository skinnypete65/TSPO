package main

import "fmt"

func main() {
	fmt.Println(checkDivisibility(15))
	fmt.Println(checkDivisibility(12))
}

func checkDivisibility(num int) string {
	if num%3 == 0 && num%5 == 0 {
		return "Делится"
	}
	return "Не делится"
}
