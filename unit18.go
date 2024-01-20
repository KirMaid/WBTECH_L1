/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика. */
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter Используем атомарные операции для обеспечения безопасности в многопоточной среде
type Counter struct {
	value int64
}

// Increment Используем atomic.AddInt64 для безопасного инкремента значения счетчика.
// Это гарантирует, что даже если несколько горутин одновременно вызывают этот метод, значение счетчика будет корректно увеличиваться.
func (c *Counter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

// Value Используем atomic.LoadInt64 для безопасного чтения значения счетчика.
// Это гарантирует, что даже если другие горутины одновременно вызывают метод Increment, значение счетчика будет корректно читаться.
func (c *Counter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	//Запускаем 1000 инкрементирующих корутин
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	//Если конфликтов не возникло, итоговое значение должно быть 1000
	fmt.Println("Финальное значение:", counter.Value())
}
