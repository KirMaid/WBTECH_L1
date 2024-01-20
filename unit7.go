/*Реализовать конкурентную запись данных в map.*/
package main

import (
	"fmt"
	"sync"
)

/*
1.Использование sync.Mutex или sync.RWMutex:
Эти структуры предоставляют механизмы блокировки для синхронизации доступа к map.
sync.Mutex используется, когда требуется взаимная исключительная блокировка, в то время как sync.RWMutex позволяет множеству горутин одновременно читать данные, но только одной горутине записывать данные.
*/
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func (sm *SafeMap) Store(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Load(key string) (int, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	val, ok := sm.m[key]
	return val, ok
}

/*
2.Использование sync.Map:
Когда ключи в map стабильны (не обновляются часто) и происходит намного больше чтений, чем записей. sync.
Map гарантирует константное время доступа к map вне зависимости от количества одновременных чтений и количества ядер.
*/

func variantMutex() {
	sm := &SafeMap{
		m: make(map[string]int),
	}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("worker-%d", id), id)
		}(i)
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		if val, ok := sm.Load(fmt.Sprintf("worker-%d", i)); ok {
			fmt.Printf("worker-%d: %d\n", i, val)
		}
	}
}

func variantSync() {
	smSync := &sync.Map{}

	var wgGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		wgGroup.Add(1)
		go func(id int) {
			defer wgGroup.Done()
			smSync.Store(fmt.Sprintf("worker-sync%d", id), id)
		}(i)
	}
	wgGroup.Wait()

	smSync.Range(func(key, value interface{}) bool {
		fmt.Printf("%v: %v\n", key, value)
		return true
	})
}

func main() {
	fmt.Println("Использование sync.Mutex или sync.RWMutex")
	variantMutex()
	fmt.Println("Использование sync.Map")
	variantSync()
}
