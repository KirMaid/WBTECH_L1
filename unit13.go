/*
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	var a, b int = 5, 10
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)

	// Используем первую функцию
	swapWithoutTempVar(a, b)
	fmt.Printf("After swapping with addition and subtraction: a = %d, b = %d\n", a, b)

	// Используем вторую функцию
	a, b = 5, 10 // Сбрасываем значения, чтобы повторить тест
	swapWithXor(&a, &b)
	fmt.Printf("After swapping with XOR: a = %d, b = %d\n", a, b)
}

func swapWithoutTempVar(a, b int) {
	b = a + b
	a = b - a
	b = b - a
}

func swapWithXor(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
