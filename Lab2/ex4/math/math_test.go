package math

import "testing"

func TestAdd(t *testing.T) {
	x := Add(1, 2, -3)
	res := 0.0
	if x != res {
		t.Errorf("Тест не пройден! Результат %f, а должен быть %f", x, res)
	}
}
func TestGetMin(t *testing.T) {
	res := GetMin(1, 2, 3)
	if res != 1 {
		t.Errorf("Тест не пройден! Результат %v, а должен быть %v", res, 1)
	}
}

func TestGetAvg(t *testing.T) {
	res := GetAvg(1, 2, 3)
	if res != 2 {
		t.Errorf("Тест не пройден! Результат %v, а должен быть %v", res, 2)
	}
}

func TestGetХ_linear(t *testing.T) {
	res := GetХ_linear( 2, 3)
	if res != -1.5 {
		t.Errorf("Тест не пройден! Результат %v, а должен быть %v", res, 1.5)
	}
}
