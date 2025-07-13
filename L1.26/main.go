/*Разработать программу, которая проверяет, что все символы в строке встречаются один раз (т.е. строка состоит из уникальных символов).

Вывод: true, если все символы уникальны, false, если есть повторения. Проверка должна быть регистронезависимой, т.е. символы в разных регистрах считать одинаковыми.

Например: "abcd" -> true, "abCdefAaf" -> false (повторяются a/A), "aabcd" -> false.

Подумайте, какой структурой данных удобно воспользоваться для проверки условия.*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var data string
	fmt.Println("введите строку для проверки уникальности символов")
	fmt.Scan(&data)

	if !unique(data) {
		fmt.Println("false - строка не состоит из уникальных символов")
		return
	}

	fmt.Println("true - строка состоит из уникальных символов")

}

func unique(s string) bool {
	// создается мапа для хранения символов (ключ - символ, значение - кол-во символов)
	mapa := make(map[string]int)

	for _, value := range s {
		mapa[strings.ToLower(string(value))]++
		if mapa[strings.ToLower(string(value))] > 1 {
			return false
		}
	}
	return true
}
