package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).

Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

Подсказка: можно использовать контекст (context.Context) или канал для оповещения о завершении.*/

/*Обоснование: При применении контекста программа завершается по истечении указанного в контексте времени,либо по по нажатию Ctrl+C (SIGINT)*/
func main() {
	//Количество воркеров
	var nWork int
	//Временной промежуток, по истечении котрого совершится отмена контекста
	var nCtx int
	fmt.Println("введите количество воркеров")
	fmt.Scan(&nWork)
	fmt.Println("введите время работы программы в секундах")
	fmt.Scan(&nCtx)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(nCtx)*time.Second)
	defer cancel()

	cIn := make(chan int)
	cCall := make(chan os.Signal, 1)
	signal.Notify(cCall, syscall.SIGINT)
	wg := sync.WaitGroup{}

	for value := range nWork {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range cIn {
				fmt.Printf("данные из воркера %d: %d\n", value, data)
				time.Sleep(500 * time.Millisecond)
			}
			fmt.Printf("вооркер %d завершил работу\n", value)
		}()
	}
	go func() {
		for range cCall {
			cancel()
		}
	}()

	for i := 1; i > 0; i++ {
		select {
		case <-ctx.Done():
			close(cIn)
			wg.Wait()
			os.Exit(0)
		default:
			cIn <- i
		}
	}
}
