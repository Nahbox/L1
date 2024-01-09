package task3

import (
	"fmt"
	"sync"
)

// Точка входа в Task3 (1-й способ решения)
func Task3_1() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создание канала для передачи результатов вычислений.
	resultChannel := make(chan int)

	var wg sync.WaitGroup

	// Итерация по числам в слайсе.
	for _, num := range numbers {
		// Увеличение счетчика горутин в WaitGroup.
		wg.Add(1)

		// Запуск горутины для вычисления квадрата числа и передачи результата в канал.
		go func(n int) {
			// Уменьшение счетчика горутин при завершении.
			defer wg.Done()
			// Вычисление квадрата числа.
			result := n * n
			// Передача результата в канал.
			resultChannel <- result
		}(num)
	}

	// Горутина для ожидания завершения всех горутин и закрытия канала.
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	// Инициализация переменной для хранения суммы результатов.
	var sum int

	// Итерация по результатам в канале и их суммирование.
	for sqr := range resultChannel {
		sum += sqr
	}

	// Вывод суммы на экран.
	fmt.Printf("%d\n", sum)
}
