package task3

import (
	"fmt"
	"sync"
)

func Task3_1() {
	numbers := []int{2, 4, 6, 8, 10}

	resultChannel := make(chan int)

	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			result := n * n
			resultChannel <- result
		}(num)
	}

	// Горутина для закрытия канала после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Собираем результаты из канала и вычисляем сумму
	var sum int
	for sqr := range resultChannel {
		sum += sqr
	}

	fmt.Printf("%d\n", sum)
}
