package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(52))
	fmt.Println(intToRoman(2024))
}

type mapping struct {
	symbol string
	val    int
}

func intToRoman(num int) string {
	nums := []mapping{
		{symbol: "M", val: 1000},
		{symbol: "CM", val: 900},
		{symbol: "D", val: 500},
		{symbol: "CD", val: 400},
		{symbol: "C", val: 100},
		{symbol: "XC", val: 90},
		{symbol: "L", val: 50},
		{symbol: "XL", val: 40},
		{symbol: "X", val: 10},
		{symbol: "IX", val: 9},
		{symbol: "V", val: 5},
		{symbol: "IV", val: 4},
		{symbol: "I", val: 1},
	}

	var ans strings.Builder
	for _, n := range nums {
		if num == 0 {
			break
		}

		cnt := num / n.val
		if cnt > 0 {
			t := strings.Repeat(n.symbol, cnt)
			ans.WriteString(t)
			num = num % n.val
		}
	}

	return ans.String()
}
