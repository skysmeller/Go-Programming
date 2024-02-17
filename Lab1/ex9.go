package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true
	//Задание.
	//1. Пояснить результаты операций
	fmt.Println("first  = ", first)       // false, по замовчуванню
	fmt.Println("second = ", second)      // false, по замовчуванню
	fmt.Println("third  = ", third)       // true, присвоєне значення
	fmt.Println("fourth = ", fourth)      // false, присвоєне значення з інверсією
	fmt.Println("fifth  = ", fifth, "\n") // true, присвоєне значення

	fmt.Println("!true  = ", !true)        // false, інверсія від true
	fmt.Println("!false = ", !false, "\n") // true, інверсія від false

	fmt.Println("true && true   = ", true && true)         // true, якщо всі умови true
	fmt.Println("true && false  = ", true && false)        // true, 1 з них false
	fmt.Println("false && false = ", false && false, "\n") // false, 2 з них false

	fmt.Println("true || true   = ", true || true)         // true, якщо хоча б 1 true
	fmt.Println("true || false  = ", true || false)        // true, якщо хоча б 1 true
	fmt.Println("false || false = ", false || false, "\n") // false, так як нема хоча б 1 true

	fmt.Println("2 < 3  = ", 2 < 3)        // true, 2 менше 3
	fmt.Println("2 > 3  = ", 2 > 3)        // false, 2 не менше 3
	fmt.Println("3 < 3  = ", 3 < 3)        // false, 3 не менше 3
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true, 3 менше або дорівнює 3
	fmt.Println("3 > 3  = ", 3 > 3)        // false, 3 не більше 3
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true, 3 більше або дорівню
	fmt.Println("2 == 3 = ", 2 == 3)       // false, 2 не дорівнює 3
	fmt.Println("3 == 3 = ", 3 == 3)       // true, 3 дорівнює 3
	fmt.Println("2 != 3 = ", 2 != 3)       // true, 2 не дорівнює 3
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false, 3 дорівнює 3


}
