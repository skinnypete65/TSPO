package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var celsiusTemp float64
	_, _ = fmt.Fscan(reader, &celsiusTemp)

	fmt.Println(celsiusTemp*1.8 + 32)
}
