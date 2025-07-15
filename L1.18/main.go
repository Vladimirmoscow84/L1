/*Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин). По завершению программы структура должна выводить итоговое значение счётчика.

Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type countWithMu struct {
	c  int64
	mu sync.Mutex
}

type countWithAtom struct {
	c int64
}

func main() {

	//количество горутин
	m := 15

	//величина счетчика
	n := 100

	wg := sync.WaitGroup{}

	// экземпляр структуры с мютексом
	count := countWithMu{}

	wg.Add(m)
	for i := 0; i < m; i++ {
		go func() {
			defer wg.Done()
			for {
				count.mu.Lock()
				if count.c < int64(n) {
					count.c++
					fmt.Printf("горутина %d получила число %d\n", i, count.c)
					count.mu.Unlock()
				} else {
					count.mu.Unlock()
					break
				}

			}
		}()

	}
	wg.Wait()
	fmt.Println("синхронизация с Мютексом")
	fmt.Printf("счетчик увеличен до %d\n", count.c)

	wg1 := sync.WaitGroup{}

	// экземпляр структуры с atomic
	count1 := countWithAtom{}

	wg1.Add(m)
	for i := 0; i < m; i++ {
		go func() {
			defer wg1.Done()
			for {
				current := atomic.LoadInt64(&count1.c)
				if current >= int64(n) {
					break
				}
				if atomic.CompareAndSwapInt64(&count1.c, current, current+1) {
					fmt.Printf("горутина %d получила число %d\n", i, count1.c)
				}
			}
		}()

	}
	wg1.Wait()
	fmt.Println("синхронизация с atomic")
	fmt.Printf("счетчик увеличен до %d\n", count1.c)
}
