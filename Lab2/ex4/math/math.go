package math

// Сумма трех чисел.
func Add(x1 float64, x2 float64, x3 float64) float64 {
	total := x1 + x2 + x3
	return total
}
func GetMin(a float64, b float64, c float64) float64 {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	} else {
		return c
	}
}

func GetAvg(a float64, b float64, c float64) float64 {
	return (a + b + c) / 3
}

func GetХ_linear(a float64, b float64) float64 {
	return (- b) / a
}
