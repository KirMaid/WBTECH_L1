/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// Рассчитываем квадрат и отправляем в канал
func calcSquare(num int, resChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	resChan <- num * num
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	//Создаём канал
	resChan := make(chan int)
	//Для корректного оповещения о завершении работы корутины создаём WaitGroup
	var wg sync.WaitGroup

	//На каждое число создаём свою корутину и передаём информацию в WaitGroup
	wg.Add(len(numbers))

	//Каждая корутина рассчитываем свой квадрат числа
	for _, num := range numbers {
		go calcSquare(num, resChan, &wg)
	}

	//Ждём завершения работы всех корутин и закрываем канал
	go func() {
		wg.Wait()
		close(resChan)
	}()

	//Выводим информацию из канала
	for res := range resChan {
		fmt.Println(res)
	}
}
