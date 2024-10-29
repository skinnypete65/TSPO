package main

import "fmt"

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)

}

func bubbleSort(arr []int) {
	fmt.Println(arr)

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			fmt.Println(arr)
		}
	}
}
