package main

import "fmt"

type FibSolver struct {
	fibNums map[int]int
}

func NewFibSolver() *FibSolver {
	return &FibSolver{
		fibNums: map[int]int{0: 0, 1: 1},
	}
}

func (f *FibSolver) Calculate(pos int) int {
	if pos < 0 {
		return 0
	}
	if val, ok := f.fibNums[pos]; ok {
		return val
	}

	f.fibNums[pos] = f.Calculate(pos-1) + f.Calculate(pos-2)
	return f.fibNums[pos]
}

func main() {
	fibSolver := NewFibSolver()

	fmt.Println(fibSolver.Calculate(10))
}

//  0  1  2  3  4  5  6  7   8   9   10  11  12   13   14   15   16   17    18    19    20    21     22
//  0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711
