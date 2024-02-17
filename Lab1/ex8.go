package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	//Задание.
	//1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести

	tmpInt := 1
	tmpString := "Hello"
	tmpFloat := 1.23
	var tmpBoolean bool

	fmt.Println(tmpInt, tmpString, tmpFloat, tmpBoolean)

}
