package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(calcUnique("saint petersburg 52 saint"))
}

type Empty struct{}

func calcUnique(str string) int {
	unique := make(map[string]Empty)
	strs := strings.Split(str, " ")

	for _, s := range strs {
		unique[s] = Empty{}
	}

	return len(unique)
}
