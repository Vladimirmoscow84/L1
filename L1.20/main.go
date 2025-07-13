/*Разработать программу, которая переворачивает порядок слов в строке.

Пример: входная строка:

«snow dog sun», выход: «sun dog snow».

Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "sun dog snow"
	fmt.Printf("исходная строка: %s\n", data)

	slice := strings.Split(data, " ")
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	fmt.Printf("строка после реверса: %s\n", strings.Join(slice, " "))

}
