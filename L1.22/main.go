/*Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

Комментарий: в Go тип int справится с такими числами, но обратите внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.*/

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b float64
	fmt.Println("введите числа")
	fmt.Scan(&a, &b)
	//переводим данные числа в big.Int
	num1 := big.NewFloat(a)
	num2 := big.NewFloat(b)

	//операция сложения
	sum := new(big.Float).Add(num1, num2)
	fmt.Printf("сложение: %s\n", sum.String())

	//опрация вычитания
	dif := new(big.Float).Sub(num1, num2)
	fmt.Printf("вычитание: %s\n", dif.String())

	//операция умножения
	mult := new(big.Float).Mul(num1, num2)
	fmt.Printf("умножение: %s\n", mult.String())

	//операция деления
	div := new(big.Float).Quo(num1, num2)
	fmt.Printf("деление: %s\n", div.String())

}
