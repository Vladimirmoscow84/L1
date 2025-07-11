package main

import "fmt"

/*Разработать программу, которая переворачивает подаваемую на вход строку.

Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.), то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).*/

// данную задачу можно решить несколькими способами, которые я представил в отдельных функциях
// 1-й способ  -  итерация по срезу рун с начала и конца с взаимной перестановкой
func firstReverse(a string) string {
	runes := []rune(a)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

// 2-й способ - итерация по строке и склейка итоговой строки с конца
func secondreverse(a string) string {
	answer := ""
	for _, v := range a {
		answer = string(v) + answer
	}
	return answer

}
func main() {
	var data string
	fmt.Println("введите строку для реверса")
	fmt.Scan(&data)
	fmt.Println(firstReverse(data))
	fmt.Println(secondreverse(data))

}
