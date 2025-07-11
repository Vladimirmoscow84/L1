package main

import (
	"fmt"
	"sort"
)

/*Реализовать алгоритм бинарного поиска встроенными методами языка. Функция должна принимать отсортированный слайс и искомый элемент, возвращать индекс элемента или -1, если элемент не найден.

Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.*/

func main() {
	slice := []int{56, 32, 3, 12, 43, 7, 89, 1, 0}
	//отсортируем слайс
	sort.Ints(slice)
	fmt.Println(slice)
	var digit int
	fmt.Println("введите искомый элемент для поиска его индекса в слайсе")
	fmt.Scan(&digit)

	//

	fmt.Printf("искомоый эемент %d находится под индексом: %d\n", digit, bSearch(slice, digit))

}
func bSearch(arr []int, n int) int {
	//определяем крание границы
	left, right := 0, len(arr)-1

	for right <= left {
		// определяем середину
		mid := right + (right+left)/2
		switch {
		case arr[mid] == n:
			return mid
		case arr[mid] > n:
			right = right - 1
		default:
			left = left + 1
		}
	}
	return -1
}
