package main

import "fmt"

func main() {
	num := 1234
	fmt.Println(sumDigits(num))
}

// Ограничение на число - число обязательно четырехзначное
func sumDigits(num int) int {
	digit1 := num % 10
	digit2 := (num / 10) % 10
	digit3 := (num / 100) % 10
	digit4 := (num / 1000) % 10

	return digit1 + digit2 + digit3 + digit4
}
