/*Реализовать бинарный поиск встроенными методами языка.*/
package main

import "fmt"

func binarySearch(arr []int, target, low, high int) int {
	if high < low {
		return -1
	}
	mid := (low + high) / 2
	if arr[mid] > target {
		return binarySearch(arr, target, low, mid-1)
	} else if arr[mid] < target {
		return binarySearch(arr, target, mid+1, high)
	} else {
		return mid
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	result := binarySearch(arr, target, 0, len(arr)-1)
	fmt.Println(result)
}
