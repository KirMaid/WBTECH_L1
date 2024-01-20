/*Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20. */
package main

import (
	"fmt"
	"math/big"
)

func main() {
	//Так как числа большие, используем big.NewInt.
	a := big.NewInt(1 << 22) // 2^22
	b := big.NewInt(1 << 21) // 2^21

	// Перемножение
	mul := new(big.Int).Mul(a, b)
	fmt.Printf("Результат умножения: %s\n", mul)

	// Деление
	div := new(big.Int).Div(a, b)
	fmt.Printf("Результат деления: %s\n", div)

	// Сложение
	add := new(big.Int).Add(a, b)
	fmt.Printf("Результат сложения: %s\n", add)

	// Вычитание
	sub := new(big.Int).Sub(a, b)
	fmt.Printf("Результат вычитания: %s\n", sub)
}
