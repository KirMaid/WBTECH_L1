/*Реализовать бинарный поиск встроенными методами языка.*/
package main

import "fmt"

// Реализовал рекурсивный вариант, так как всегда нравился больше
func binarySearch(arr []int, target, low, high int) int {
	//Это проверка базового случая для рекурсии.Если элемент не найден, выводится -1.
	if high < low {
		return -1
	}
	//Вычисляем середину массива(индекс)
	mid := (low + high) / 2

	//В зависимости от того, больше искомый элемент элемента с середины массива или меньше, рекурсивно продолжаем поиск в разных подмассивах
	if arr[mid] > target {
		return binarySearch(arr, target, low, mid-1)
	} else if arr[mid] < target {
		return binarySearch(arr, target, mid+1, high)
	} else {
		//Возвращаем результат поиска
		return mid
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	result := binarySearch(arr, target, 0, len(arr)-1)
	fmt.Println(result) //Если элемент не найден, выводится -1.
}
