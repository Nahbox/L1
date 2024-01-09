package task2

import (
	"fmt"
	"sync"
)

// Точка входа в Task2 (2-й способ решения)
func Task2_2() {
	numbers := []int{2, 4, 6, 8, 10}

	// Создание канала для передачи результатов вычислений.
	results := make(chan int, len(numbers))

	var wg sync.WaitGroup

	// Итерация по числам в слайсе.
	for _, val := range numbers {
		// Увеличение счетчика горутин в WaitGroup.
		wg.Add(1)

		// Запуск горутины для вычисления квадрата числа и передачи результата в канал.
		go powCalculation(val, results, &wg)
	}

	// Горутина для ожидания завершения всех горутин и закрытия канала.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Итерация по результатам в канале и вывод их на экран.
	for res := range results {
		fmt.Println(res)
	}
}

// Функция powCalculation вычисляет квадрат числа и передает результат в канал.
func powCalculation(number int, results chan<- int, wg *sync.WaitGroup) {
	// Уменьшение счетчика горутин при завершении.
	defer wg.Done()
	// Вычисление квадрата числа.
	square := number * number
	// Передача результата в канал.
	results <- square
}
