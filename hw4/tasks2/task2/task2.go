package main

import "fmt"

func main() {
	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(2015))
}

func isLeapYear(year int) string {
	if year%4 == 0 {
		return "Високосный"
	}
	return "Невисокосный"
}
