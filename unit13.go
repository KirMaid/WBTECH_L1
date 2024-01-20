/*
Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	var a, b int = 5, 10
	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Используем первую функцию
	swap(&a, &b)
	fmt.Printf("После обмена(способ с вычитанием и сложением): a = %d, b = %d\n", a, b)

	// Используем вторую функцию
	a, b = 5, 10 // Сбрасываем значения, чтобы повторить тест
	swapWithXor(&a, &b)
	fmt.Printf("После обмена (способ с XOR): a = %d, b = %d\n", a, b)
}

func swap(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swapWithXor(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
