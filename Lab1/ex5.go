package main

import (
	"fmt"
)

func main() {
	fmt.Println("Синонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	//Задание.
	//1. Определить разрядность ОС
	const is64Bit = uint64(^uintptr(0)) == ^uint64(0)
	if is64Bit {
		fmt.Println("ОС - 64 біт")
	} else {
		fmt.Println("ОС - 32 біт")
	}
}
