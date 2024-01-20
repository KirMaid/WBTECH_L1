/*Реализовать собственную функцию sleep.*/
package main

import (
	"fmt"
	"time"
)

// Использовал функцию NewTimer, которая возвращает таймер, который закроется через указанный промежуток времени.
func mySleep(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
}

func main() {
	fmt.Println("Начальное время:", time.Now().Format(time.Stamp))
	mySleep(2 * time.Second)
	fmt.Println("Конечное время:", time.Now().Format(time.Stamp))
}
