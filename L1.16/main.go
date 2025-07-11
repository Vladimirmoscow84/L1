package main

import "fmt"

/*Реализовать алгоритм быстрой сортировки массива встроенными средствами языка. Можно использовать рекурсию.

Подсказка: напишите функцию quickSort([]int) []int которая сортирует срез целых чисел. Для выбора опорного элемента можно взять середину или первый элемент.*/

func main() {
	arr := []int{3, 45, 12, 4, 8, 88, 92, 35, 16}
	sortedArr := quickSort(arr)
	fmt.Printf("до сортировки: %v\n", arr)
	fmt.Printf("после сортировки: %v\n", sortedArr)

}

func quickSort(nonSortedArr []int) []int {
	if len(nonSortedArr) < 2 {
		return nonSortedArr
	}
	//piv - опорный элемент
	piv := nonSortedArr[0]
	//less - слайс со значениями меньше piv
	//greater - слайс созначениями больше Piv
	var less, greater []int

	for _, value := range nonSortedArr[1:] {
		if value <= piv {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}
	answer := append(quickSort(less), piv)
	answer = append(answer, quickSort(greater)...)

	return answer
}
