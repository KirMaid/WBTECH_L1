/*Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/
package main

import (
	"fmt"
	_ "math/rand"
	_ "time"
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var pi int
		arr, pi = partition(arr, low, high)
		arr = quickSort(arr, low, pi-1)
		arr = quickSort(arr, pi+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func main() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("Изначальный массив: ", arr)
	quickSortStart(arr)
	fmt.Println("Отсортированный массив: ", arr)
}
