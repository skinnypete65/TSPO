package main

import "fmt"

func main() {
	fmt.Println(isNumPalindrome(1221))
	fmt.Println(isNumPalindrome(1231))
	fmt.Println(isNumPalindrome(121))
}

func isNumPalindrome(num int) bool {
	if num/10 == 0 {
		return true
	}

	reversed, size := reverseNum(num)
	if reversed == -1 {
		return false
	}

	for i := 0; i < size/2; i++ {
		digit1, digit2 := num%10, reversed%10

		if digit1 != digit2 {
			return false
		}

		num /= 10
		reversed /= 10
	}

	return true
}

func reverseNum(num int) (int, int) {
	newNum := 0
	size := 0

	for num > 0 {
		old := newNum
		newNum *= 10
		newNum += num % 10

		if newNum/10 != old {
			return -1, 0
		}
		num /= 10
		size++
	}

	return newNum, size
}
