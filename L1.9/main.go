package main

import "fmt"

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2. После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка. Убедитесь, что чтение из второго канала корректно завершается.*/

func main() {
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	chIn := generator(arr...)
	chOut := multiply(chIn)

	for answer := range chOut {
		fmt.Println(answer)
	}

}
func generator(arr ...int) chan int {
	chOut := make(chan int)
	go func() {
		defer close(chOut)
		for _, data := range arr {
			chOut <- data

		}
	}()
	return chOut
}
func multiply(chIn chan int) chan int {
	chOut := make(chan int)
	go func() {
		defer close(chOut)
		for data := range chIn {
			chOut <- data * 2
		}
	}()
	return chOut
}
