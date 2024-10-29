package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func startGame(attemptsCnt int) {
	randNum := rand.Intn(99) + 1
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < attemptsCnt; i++ {
		var userNum int
		fmt.Fscan(reader, &userNum)

		if userNum == randNum {
			fmt.Printf("Вы угадали число! Это число - %d\n", userNum)
			return
		}

		if userNum > randNum {
			fmt.Println("Ваше число больше")
		} else {
			fmt.Println("Ваше число меньше")
		}
	}

	fmt.Printf("Вы не угадали число :( Это число - %d\n", randNum)
}

func main() {
	startGame(10)
}
