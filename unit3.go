/*Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2 в квадрате + +3 в квадрате +4 в квадрате….) с использованием конкурентных вычислений.*/

package main

import (
	"fmt"
	"sync"
)

// Рассчитываем квадрат и отправляем в канал
func calc(num int, resChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	resChan <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//Для корректного оповещения о завершении работы корутины создаём WaitGroup
	var wg sync.WaitGroup
	results := make(chan int)

	//На каждое число создаём свою корутину и передаём информацию в WaitGroup
	wg.Add(len(numbers))
	for _, num := range numbers {
		go calc(num, results, &wg)
	}

	//Ждём завершения работы всех корутин и закрываем канал
	go func() {
		wg.Wait()
		close(results)
	}()

	//Считаем и выводим итоговую сумму
	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println(sum)
}
