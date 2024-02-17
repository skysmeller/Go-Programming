// main
package main

import (
	"fmt"

	"./math"
)

func main() {

	a := 1.0
	b := 2.5
	c := 5.3
	fmt.Printf("min: %v \n", math.GetMin(a, b, c))
	fmt.Printf("avg: %v \n", math.GetAvg(a, b, c))
	fmt.Printf("x: %v \n", math.GetХ_linear(b,c))
}
