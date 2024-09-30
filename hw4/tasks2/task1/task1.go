package main

import "fmt"

func main() {
	fmt.Println(isEven(4))
	fmt.Println(isEven(5))
}

func isEven(num int) string {
	if num%2 == 0 {
		return "Четное"
	}
	return "Нечетное"
}
