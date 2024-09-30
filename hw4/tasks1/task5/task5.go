package main

import (
	"fmt"
	"math"
)

func main() {
	dist := calcDist(1, 1, 4, 5)
	fmt.Println(dist)
}

func calcDist(x1, y1, x2, y2 float64) float64 {
	xDiff := math.Pow(x1-x2, 2)
	yDiff := math.Pow(y1-y2, 2)

	return math.Sqrt(xDiff + yDiff)
}
