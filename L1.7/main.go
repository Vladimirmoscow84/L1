package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*Реализовать безопасную для конкуренции запись данных в структуру map.

Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

Проверьте работу кода на гонки (util go run -race).*/

type data struct {
	mapa map[int]int
	mu   sync.Mutex
}

func main() {

	m := data{
		mapa: make(map[int]int),
		mu:   sync.Mutex{},
	}
	wg := sync.WaitGroup{}

	for i := 1; i <= 30; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			m.mu.Lock()
			m.mapa[v] = rand.Intn(1000)
			m.mu.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println(m.mapa)
}
