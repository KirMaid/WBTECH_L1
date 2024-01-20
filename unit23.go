/*Удалить i-ый элемент из слайса. */
package main

import "fmt"

// Решение - использовать функцию append вместе с срезкой слайса.
func deleteAtIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	numbers := []int{10, 20, 30, 40, 50, 90, 60}
	index := 3

	fmt.Println("Оригинальный слайс:", numbers)
	fmt.Printf("Элемент %d будет удалён.\n", numbers[index])
	numbers = deleteAtIndex(numbers, index)
	fmt.Println("Слайс после удаления элемента:", numbers)
}
