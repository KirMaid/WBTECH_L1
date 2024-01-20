package main

import (
	"fmt"
	"os"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.
*/

// Каждую секунду шлёт значение в канал
func sender(c chan int) {
	for i := 0; ; i++ {
		c <- i
		time.Sleep(time.Second)
	}
}

// Функция для считывания значений из канала
func reader(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	var programTime int
	fmt.Print("Введите время выполнения программы, в секундах:")
	_, err := fmt.Fscan(os.Stdin, &programTime)
	if err != nil {
		fmt.Print("Неверный формат ввода")
		return
	}
	c := make(chan int)
	go sender(c)
	go reader(c)
	//Завершаем работы программы после истечения времени
	time.Sleep(time.Duration(programTime) * time.Second)
}
