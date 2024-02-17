
package random

func GetRandIntNums(x int, a int, c int, m int, start int, end int, length int) []int {

	numbers := make([]int, length)
	mod := end-start + 1

	for i := 0; i < length; i++ {
		x = (a*x + c) % m
		numbers[i] = x % mod
	}

	return numbers
}


func GetRangFloatNums(x int, a int, c int, m int, startRange int, endRange int, length int) []float64 {

	numbers := make([]float64, length)

	for i := 0; i < length; i++ {
		x = (a*x + c) % m
		numbers[i] = float64(x) / (float64(m) / float64(startRange+x%endRange))
	}

	return numbers
}


