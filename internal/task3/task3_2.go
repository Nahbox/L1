package task3

import (
	"fmt"
	"sync"
)

func Task3_2() {
	numbers := []int{2, 4, 6, 8, 10}

	var mu sync.Mutex
	// Общая переменная для хранения суммы квадратов
	var sum int

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sqr := n * n

			// Захватываем мьютекс для безопасного обновления общей переменной
			mu.Lock()
			sum += sqr
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Printf("%d\n", sum)
}
