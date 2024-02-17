package random

import (
	"math"
	"testing"

)

func TestGetRandIntNums(t *testing.T) {
	a := 16807
	c := 0
	m := int(math.Pow(2, 31)-1)
	x := 1
	numRange := [2]int{0, 200}
	length := 3000


	intNumbers := GetRandIntNums(x, a, c, m, numRange[0], numRange[1], length)

	res := 113

	if intNumbers[5] != res {
		t.Errorf("Тест не пройден! Результат %v, а должен быть %v", intNumbers[5], res)
	}
}
func TestGetRangFloatNums(t *testing.T) {

	a := 16807
	c := 0
	m := int(math.Pow(2, 31)-1)
	x := 1
	numRange := [2]int{0, 200}
	length := 3000


	intNumbers := GetRangFloatNums(x, a, c, m, numRange[0], numRange[1], length)

	res := 6.445351619015146

	if intNumbers[1] != res {
		t.Errorf("Тест не пройден! Результат %v, а должен быть %v", intNumbers[1], res)
	}
}