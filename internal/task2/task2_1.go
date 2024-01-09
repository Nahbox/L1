package task2

import (
	"fmt"
	"sync"
)

// Точка входа в Task2 (1-й способ решения)
func Task2_1() {
	numbers := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	// Итерация по числам в слайсе.
	for _, val := range numbers {
		// Увеличение счетчика горутин в WaitGroup.
		wg.Add(1)

		// Запуск горутины для вывода квадрата числа.
		go func(val int) {
			// Уменьшение счетчика горутин при завершении.
			defer wg.Done()
			// Вывод квадрата числа.
			fmt.Println(val * val)
		}(val)
	}

	// Ожидание завершения всех горутин.
	wg.Wait()
}
