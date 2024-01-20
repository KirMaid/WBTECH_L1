/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/
package main

import "fmt"

func main() {
	strSequence := []string{"cat", "cat", "dog", "cat", "tree"}
	strSet := make(map[string]struct{})

	for _, str := range strSequence {
		strSet[str] = struct{}{}
	}

	fmt.Println(strSet)
}
