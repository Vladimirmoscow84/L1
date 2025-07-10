package main

import "fmt"

/*Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}). Типы, которые нужно распознавать: int, string, bool, chan (канал).

Подсказка: оператор типа switch v.(type) поможет в решении.*/

func main() {

	arr := make([]interface{}, 5)

	arr = []interface{}{5, "WB Tech", true, make(chan interface{}), 2.2}

	for _, v := range arr {
		switch v.(type) {
		case int:
			fmt.Printf("тип переменной v - %T\n", v)
		case string:
			fmt.Printf("тип переменной v - %T\n", v)
		case bool:
			fmt.Printf("тип переменной v - %T\n", v)
		case chan interface{}:
			fmt.Println("тип переменной v - chan")
			//fmt.Printf("тип переменной v - %T\n", v)
		default:
			fmt.Printf("неизвестный тип переменной")
		}
	}

}
