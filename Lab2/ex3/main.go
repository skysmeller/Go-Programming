// main
package main

import (
	"fmt"

	mymath "./math"
)

func main() {
	a := 1.0
	b := 2.5
	c := 5.3
	fmt.Printf("min: %v \n", mymath.GetMin(a, b, c))
	fmt.Printf("avg: %v \n", mymath.GetAvg(a, b, c))
	fmt.Printf("x: %v \n", mymath.GetХ_linear(b,c))
}
