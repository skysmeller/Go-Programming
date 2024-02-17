// main
package main

import "fmt"

func main() {
	a := 1.0
	b := 2.5
	c := 5.3
	fmt.Printf("min: %v \n", getMin(a, b, c))
	fmt.Printf("avg: %v \n", getAvg(a, b, c))
	fmt.Printf("x: %v \n", getХ_linear(b,c))
}

func getMin(a float64, b float64, c float64) float64 {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func getAvg(a float64, b float64, c float64) float64 {
	return (a + b + c) / 3
}

func getХ_linear(a float64, b float64) float64 {
	return (- b) / a
}
