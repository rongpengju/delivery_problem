package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Round(1.01, 0))
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
