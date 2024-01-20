/*Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0. */
package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	n |= 1 << pos
	return n
}

func clearBit(n int64, pos uint) int64 {
	n &^= 1 << pos
	return n
}

func main() {
	var num int64 = 10 // 1010 в двоичном формате
	fmt.Printf("Оригинальное число: %064b\n", num)

	// Установить 2-й бит в 1
	num = setBit(num, 3)
	fmt.Printf("После установки единицы в четвёртый разряд: %064b\n", num)

	// Сбросить 2-й бит в 0
	num = clearBit(num, 3)
	fmt.Printf("После установки нуля в четвёртый разряд: %064b\n", num)
}
