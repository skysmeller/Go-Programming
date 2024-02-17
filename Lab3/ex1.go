package main

import (
	"./random"
	"fmt"
	"math"
)

func main() {
	fmt.Println("Завдання №1")

	a := 16807
	c := 0
	m := int(math.Pow(2, 31)-1)
	x := 1
	numRange := [2]int{0, 200}
	length := 3000

	intNumbers := random.GetRandIntNums(x, a, c, m, numRange[0], numRange[1], length)
	fmt.Println(intNumbers)

	stats(intNumbers, numRange[0],numRange[1]);



	fmt.Println("Завдання №2")
	floatNumbers := random.GetRangFloatNums(x, a, c, m, numRange[0], numRange[1], length)

	fmt.Println(floatNumbers)
}

func stats(numbers []int, start int, end int)  {

	mx := 0.0
	d := 0.0

	for i := start; i <= end; i++ {

		countNum := 0

		for j := 0; j < len(numbers); j++ {
			if i == numbers[j] {
				countNum++

			}
		}
		p := float64(countNum) / float64(len(numbers))
		mx += float64(i) * p
		d += math.Pow(float64(i)-mx, 2) * p

		fmt.Printf("Число: %v\n",i)
		fmt.Printf("\tЧастота: \t\t%v,\n\t Ймовірність: \t%v\n", countNum, p)

	}
	fmt.Printf("M(X) - %f\n", mx)
	fmt.Printf("D(X)  - %f\n", d)
	fmt.Printf("Сер. відхилення - %f\n",  math.Sqrt(d))

}