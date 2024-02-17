package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"
	tmp := 'Ї'
	fmt.Printf("Code '%c' - %d\n", tmp, int16(tmp))

	//rune як псевдонім для типу int32

}
