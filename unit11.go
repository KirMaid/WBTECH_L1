// Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)

// Функция проходит циклом по первому массиву и находим элементы с таким же ключом во втором массиве, после чего сохраняет результат в отдельный массив
func intersection(set1, set2 map[int]bool) []int {
	var intersection []int
	for key := range set1 {
		if set2[key] {
			intersection = append(intersection, key)
		}
	}

	return intersection
}

func main() {
	set1 := map[int]bool{2: true, 3: true, 5: true, 6: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}
	fmt.Println(intersection(set1, set2))
}
