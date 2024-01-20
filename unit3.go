package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, results chan<- int) {
	results <- num * num
	wg.Done()
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	results := make(chan int)

	wg.Add(len(numbers))
	for _, num := range numbers {
		go square(num, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println(sum)
}
