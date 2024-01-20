/*Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode. */
package main

import "fmt"

func reverse(s string) string {
	//Преобразуем строку в слайс по символам
	r := []rune(s)
	//Итерационно меняем попарно символы местами(первые и последние)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	s := "главрыба"
	fmt.Println(reverse(s))
}
