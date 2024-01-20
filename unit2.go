/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/
package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа и отправки результата в канал.
func squareWorker(numbers <-chan int, squares chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range numbers {
		squares <- number * number // отправляем квадрат числа в канал
	}
}

// Функция для подсчета итоговой суммы из канала квадратов.
func sumWorker(squares <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for square := range squares {
		sum += square
	}
	fmt.Printf("Сумма квадратов: %d\n", sum)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10} // последовательность чисел
	numWorker := 5                   // количество горутин-воркеров

	var wg sync.WaitGroup
	numberChan := make(chan int, len(numbers))
	squareChan := make(chan int, numWorker)

	// Запускаем воркеры для вычисления квадратов
	wg.Add(numWorker)
	for i := 0; i < numWorker; i++ {
		go squareWorker(numberChan, squareChan, &wg)
	}

	// Запускаем воркера для подсчета суммы
	wg.Add(1)
	go sumWorker(squareChan, &wg)

	// Отправляем числа в канал
	for _, number := range numbers {
		numberChan <- number
	}
	close(numberChan) // закрываем канал после отправки всех чисел

	// Ждем завершения всех горутин
	wg.Wait()
	close(squareChan) // закрываем канал после завершения работы воркеров
}
