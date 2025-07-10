package main

import (
	"fmt"
)

/*Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.*/

func main() {

	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	mapa := make(map[string]bool)
	slice := make([]string, 0)
	for _, v := range arr {
		mapa[v] = true
	}
	for v := range mapa {
		slice = append(slice, v)
	}
	//можно отсортировать
	//sort.Strings(slice)
	fmt.Printf("Множество = %v", slice)

}
