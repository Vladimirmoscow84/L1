package main

import (
	"fmt"
	"time"
)

/*Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

Подсказка: используйте time.After или таймер для ограничения времени работы.*/

func main() {
	ch := make(chan int)
	st := time.Now()

	go func() {
		for i := 1; i > 0; i++ {
			select {
			//проверка состояния канала
			case _, ok := <-ch:
				if !ok {
					fmt.Println("Писатель вскрылся")
					return
				}
			case ch <- i:
			}
		}
	}()
	//запуск в отдельной горутине алгоритма для закрытия канала
	go func() {
		<-time.After(5 * time.Second)
		close(ch)
	}()

	for value := range ch {
		fmt.Printf("Данные из канала: %d\n", value)
	}
	fmt.Printf("Время истекло, программа закрывается, время работы: %v\n", time.Since(st))
}
