package main

import (
	"fmt"
)

/*Дана переменная типа int64. Разработать программу, которая устанавливает i-й бит этого числа в 1 или 0.

Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).

Подсказка: используйте битовые операции (|, &^).*/

func main() {
	var number int64
	fmt.Println("введите число")
	fmt.Scan(&number)
	strnumber := fmt.Sprintf("%b", number)
	fmt.Printf("числом %d в бинарном представлении является %s\n", number, strnumber)
	var ind int
	fmt.Println("введите индекс i-того бита для установки (начиная с 0)")
	fmt.Scan(&ind)

	if ind > len(strnumber)-1 || ind < 0 {
		fmt.Printf("такого индекса бита в бинарном представлении числа %d нет. Выход из программы", number)
		return
	}
	var bit int
	fmt.Printf("введите значение бита: %d или %d\n", 0, 1)
	fmt.Scan(&bit)
	if bit != 1 && bit != 0 {
		fmt.Println("неверное значение бита. Выход из программы")
		return
	}

	var mask, result int64

	mask = 1 << ind
	if bit == 1 {
		result = number | mask
	} else {
		result = number &^ mask
	}

	fmt.Printf("Измененное число: %d (%b)", result, result)

}
