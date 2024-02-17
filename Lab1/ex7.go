package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	//Задание.
	//1. Создайте 2 переменные  разных типов. Выпоните арифметические операции. Результат вывести

	var tmp1 uint8 = 127
	var tmp2 int64 = 1024
	fmt.Println("a * b = ",tmp1*uint8(tmp2))
	fmt.Println("a * b = ", int64(tmp1)*tmp2)
	fmt.Println("a - b = ", tmp1-uint8(tmp2))
	fmt.Println("a - b = ", int64(tmp1)-tmp2)


}
