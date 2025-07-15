/*Реализовать собственную функцию sleep(duration) аналогично встроенной функции time.Sleep, которая приостанавливает выполнение текущей горутины.

Важно: в отличии от настоящей time.Sleep, ваша функция должна именно блокировать выполнение (например, через таймер или цикл), а не просто вызывать time.Sleep :) — это упражнение.

Можно использовать канал + горутину, или цикл на проверку времени (не лучший способ, но для обучения).*/

package main

import (
	"fmt"
	"time"
)

func timeToSleep(d time.Duration) {
	ch := make(chan struct{})

	go func() {
		time.Sleep(d)
		close(ch)

	}()
	<-ch

}

func main() {
	//количество секунд
	var n int
	fmt.Println("введите количество секунд")
	fmt.Scan(&n)
	fmt.Println("начало работы функции")
	now := time.Now()
	timeToSleep(time.Duration(int64(n)) * time.Second)
	fmt.Printf("время работы функции составило:%v", time.Since(now))

}
