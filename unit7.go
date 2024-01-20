/*Реализовать конкурентную запись данных в map.*/
package main

import "sync"

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
