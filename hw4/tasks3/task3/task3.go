package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	reverseSlice(sl)
	fmt.Println(sl)
}

func reverseSlice(sl []int) {
	for l := 0; l < len(sl)/2; l++ {
		r := len(sl) - l - 1

		sl[l], sl[r] = sl[r], sl[l]
	}
}
