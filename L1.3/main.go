package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Реализовать постоянную запись данных в канал (в главной горутине).

Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.

Программа должна принимать параметром количество воркеров и при старте создавать указанное число горутин-воркеров.*/

func main() {
	var numberWorkers int
	fmt.Println("Input number of workers")
	fmt.Scan(&numberWorkers)
	wg := sync.WaitGroup{}

	// канал для записи данных в главной горутине
	cIn := make(chan int)

	//канал для получения сигнала на закрытие канала в главной горутине
	cCall := make(chan os.Signal, 1)

	signal.Notify(cCall, syscall.SIGINT)

	for value := range numberWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range cIn {
				fmt.Printf("worker %d: %d\n", value, data)
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Printf("worker %d has closed\n", value)

		}()

	}

	for i := 1; i > 0; i++ {
		select {
		case <-cCall:
			close(cIn)
			wg.Wait()
			fmt.Println("end of work")
			os.Exit(0)
		default:
			cIn <- i
		}
	}

}
