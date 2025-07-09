package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

/*Реализовать все возможные способы остановки выполнения горутины.

Классические подходы: выход по условию, через канал уведомления, через контекст, прекращение работы runtime.Goexit() и др.

Продемонстрируйте каждый способ в отдельном фрагменте кода.*/

func main() {

	//выход из горутины через канал уведомленя
	chOut := make(chan struct{})
	go exitWithChannel(chOut)
	close(chOut)

	//выход из горутины при отмене контекста
	ctx, cancel := context.WithCancel(context.Background())
	go exitWithCtxCancel(ctx)
	cancel()

	//выход из горутины через контекст по истечении заданного интервала времени
	deadLine := time.Now().Add(500 * time.Millisecond)
	ctx, cancel = context.WithDeadline(context.Background(), deadLine)
	go exitWithCtxDeadline(ctx)
	defer cancel()

	// выход из горутины через указанный промежуток времени (time.After)
	go exitWithTimeAfter()

	//выход из горутины по времени при реализации таймера (NewTimer)
	go exitWithTimer()

	//выход из горутины через спеиальную функцию runtime.Goexit()
	go exitWithRuntime()

	time.Sleep(2 * time.Second)

}

func exitWithChannel(ch chan struct{}) {

	for {
		select {
		case <-ch:
			fmt.Println("горутина завершила работу по сигналу канала уведолмения")
			return
		}
	}
}

func exitWithCtxCancel(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("горутина заершила работу при отмене контекста")
			return
		}
	}
}

func exitWithCtxDeadline(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("горутина завершила работу по истечении заданного интервала в контексте")
			return
		}
	}
}

func exitWithTimeAfter() {
	timeToExit := time.After(1 * time.Second)
	for {
		select {
		case <-timeToExit:
			fmt.Println("горутина завершила работу через указанный промежуток времени")
			return
		}
	}
}
func exitWithTimer() {
	t := time.NewTimer(600 * time.Millisecond)
	for {
		select {
		case <-t.C:
			fmt.Println("горутина завершила работу по истечении таймера")
			return
		}
	}
}

func exitWithRuntime() {
	time.Sleep(800 * time.Millisecond)
	fmt.Println(`горутина завершила работу через функцию "runtime.Goexit()"`)
	runtime.Goexit()
}
