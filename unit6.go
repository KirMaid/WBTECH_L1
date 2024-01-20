package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// 1.Остановка по состоянию канала
func workerCanal(stop <-chan bool) {
	for {
		select {
		default:
			fmt.Println("Работаю(Канал)...")
			time.Sleep(1 * time.Second)
		case <-stop:
			fmt.Println("Остановлен по сигналу из канала")
			return
		}
	}
}

// 2.Использование контекста
func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Остановлен по контексту")
			return
		default:
			fmt.Println("Работаю(контекст)...")
			time.Sleep(1 * time.Second)
		}
	}
}

// 3.Остановка работы корутины по завершению основного потока программы
func workerStandart() {
	for {
		fmt.Println("Работаю...")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// 1.Остановка по состоянию канала
	stop := make(chan bool)
	go workerCanal(stop)
	time.Sleep(3 * time.Second)
	stop <- true
	time.Sleep(1 * time.Second)

	//2.Использование контекста
	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)

	// 3.Остановка работы корутины по завершению основного потока программы
	go workerStandart()
	time.Sleep(3 * time.Second)

	fmt.Println("Программа завершена")
}
