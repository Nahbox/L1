package task2

import (
	"fmt"
	"sync"
)

func Task2_2() {
	numbers := []int{2, 4, 6, 8, 10}

	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	for _, val := range numbers {
		wg.Add(1)
		go powCalculation(val, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}

func powCalculation(number int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	square := number * number
	results <- square
}
