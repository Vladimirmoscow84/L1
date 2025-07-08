package main

import (
	"fmt"
	"sync"
	"time"
)

/*Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.*/

func main() {
	var numberWorkers int
	fmt.Println("Введите количество воркеров")
	fmt.Scan(&numberWorkers)
	wg := sync.WaitGroup{}

	cIn := make(chan int)

	for value := range numberWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range cIn {
				fmt.Printf("worker %d: %d\n", value, data)
			}
			fmt.Printf("worker %d has closed\n", value)

		}()

	}

	for i := 1; i < 15; i++ {
		cIn <- i
		time.Sleep(1 * time.Second)
	}
	fmt.Println("The work has done")
	close(cIn)
	wg.Wait()

}
