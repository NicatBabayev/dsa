package main

import (
	"fmt"
	"math"
)

func sumOfGeometricSeries(a, r, n float64) float64 {
	res := (a * (1 - math.Pow(r, n))) / (1 - r)
	return res
}

func main() {
	fmt.Println(sumOfGeometricSeries(2, 2, 15))

}
