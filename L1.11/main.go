package main

import "fmt"

/*Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

Пример:
A = {1,2,3}
B = {2,3,4}
Пересечение = {2,3}*/

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 3, 4}

	mapa := make(map[int]bool)
	slice := make([]int, 0)

	for _, v := range arr1 {
		mapa[v] = true
	}

	for _, v := range arr2 {
		if mapa[v] {
			slice = append(slice, v)
		}
	}
	fmt.Printf("Пересечение = %v", slice)
}
