package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	fmt.Scan(&r)
	area := CircleArea(r)
	fmt.Printf("圆的面积为: %.2f\n", area)
}

func CircleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
