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

	fmt.Printf("строка после реверса: %s\n", reverse(data))

}
func reverse(s string) string {
	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	s = strings.Join(words, " ")
	return s

}
