package task2

import (
	"fmt"
	"sync"
)

func Task2_1() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	for _, val := range numbers {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Println(val * val)
		}(val)
	}

	wg.Wait()
}
