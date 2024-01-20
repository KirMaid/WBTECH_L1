/*Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.*/
package main

import (
	"fmt"
	"unicode"
)

// Используем символы как ключи в map.В случае, если такой ключ уже существует, значит в строке не ве символы уникальные.
func uniqueChars(str string) bool {
	seen := make(map[rune]struct{})
	for _, char := range str {
		//Приводим всё в нижний регистр
		char = unicode.ToLower(char)
		if _, ok := seen[char]; ok {
			return false
		}
		seen[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(uniqueChars("abcd"))      //True
	fmt.Println(uniqueChars("abCdefAaf")) //False
	fmt.Println(uniqueChars("aabcd"))     //False
}
