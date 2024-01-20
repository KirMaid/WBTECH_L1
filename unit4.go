/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Job struct {
	data int
}

func testWorker(id int, jobs <-chan Job, stop <-chan bool, done chan<- bool) {
	for {
		select {
		case <-stop:
			fmt.Printf("Worker %d получил стоп сигнал\n", id)
			done <- true
			return
		case job := <-jobs:
			fmt.Printf("Worker %d обрабатывает data %d\n", id, job.data)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	//Кол-во воркеров вводим с консоли
	var numWorkers int
	fmt.Print("Введите количество воркеров:")
	_, err := fmt.Fscan(os.Stdin, &numWorkers)
	if err != nil {
		fmt.Print("Неверный формат ввода")
		return
	}

	//Создаём пул задач
	jobs := make(chan Job, numWorkers)

	//Стоп сигнал присылаем в каждый воркер по отдельному каналу. Это необходимо для безопасного завершения работы корутины.
	//В ином случае корутина попытается записать в закрытый канал и мы получим panic.
	//Поэтому программа завершается не сразу после нажатия CTRL C, а после завершения работы всех корутин.

	stop := make(chan bool)
	done := make(chan bool, numWorkers)

	// Запускаем воркеров
	for w := 1; w <= numWorkers; w++ {
		go testWorker(w, jobs, stop, done)
	}

	// Отправляем задачи воркерам
	go func() {
		for d := 1; ; d++ {
			jobs <- Job{data: d}
		}
	}()

	// Создаём сочетание клавиш для завершения работы
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Ожидаем сигнал завершения работы
	<-sigs
	fmt.Println("\nПолучен сигнал завершения работы...")
	for i := 0; i < numWorkers; i++ {
		stop <- true
	}

	// Ожидаем завершения всех воркеров
	for i := 0; i < cap(done); i++ {
		<-done
	}
}
