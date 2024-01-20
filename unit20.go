/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	//Преобразуем строку в слайс по словам
	words := strings.Fields(s)
	//Итерационно меняем попарно слова местами(первые и последние)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	//Склеиваем обратно
	return strings.Join(words, " ")
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
