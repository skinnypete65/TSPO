package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"Hello", "World", "!"}
	fmt.Println(joinStrings(s))
}

func joinStrings(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(strs[0])

	for i := 1; i < len(strs); i++ {
		sb.WriteString(" ")
		sb.WriteString(strs[i])
	}

	return sb.String()
}
