package main

import "fmt"

func main() {
	fmt.Println(reverseNum(12345))
}

func reverseNum(num int) int {
	newNum := 0

	for num > 0 {
		old := newNum
		newNum *= 10
		newNum += num % 10
		
		if newNum/10 != old {
			return -1
		}
		num /= 10
	}

	return newNum
}
