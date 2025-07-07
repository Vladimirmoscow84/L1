package main

import (
	"fmt"
	"sync"
)

/*Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.*/

func readFromWorkers(cOut chan int, wg *sync.WaitGroup) int {
	value := <-cOut
	wg.Done()
	return value
}

func main() {
	var numberWorkers int
	fmt.Println("Введите количество воркеров")
	fmt.Scan(&numberWorkers)
	cIn := make(chan int)
	for i := 0; i < 5; i++ {
		cIn <- i
	}

	wg := sync.WaitGroup{}
	for range numberWorkers {
		wg.Add(1)
		go readFromWorkers(cIn, &wg)
	}

	wg.Wait()
	close(cIn)

}
