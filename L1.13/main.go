package main

import "fmt"

/*Поменять местами два числа без использования временной переменной.

Подсказка: примените сложение/вычитание или XOR-обмен.*/

func main() {
	fmt.Println("1 метод перестановки:")
	a := 5
	b := 10
	fmt.Printf("a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)

	fmt.Println("2 метод сложения-вычитания:")
	a = 5
	b = 10
	fmt.Printf("a=%d, b=%d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%d, b=%d\n", a, b)

	fmt.Println("3 XOR метод:")
	a = 5
	b = 10
	fmt.Printf("a=%d, b=%d\n", a, b)
	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("a=%d, b=%d\n", a, b)

}
