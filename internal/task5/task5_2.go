package task5

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Функция Task5_2 выполняет задачу с использованием context.WithTimeout.
func Task5_2() {
	// Переменная для хранения количества секунд.
	var secCount int
	fmt.Print("Количество секунд: ")
	fmt.Scan(&secCount)
	// Запись текущего времени.
	start := time.Now()

	// Создание контекста с таймаутом.
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(secCount)*time.Second)
	// Создание канала для передачи данных.
	data := make(chan any)

	// Запуск горутины для писателя с использованием контекста.
	go writer2(data, ctx)
	// Запуск горутины для воркера с использованием контекста.
	go worker2(data, ctx)

	// Ожидание указанного интервала времени.
	time.Sleep(time.Duration(secCount) * time.Second)
	// Вывод времени выполнения.
	fmt.Println("\nВремя выполнения: ", time.Since(start))
}

// Функция writer2 генерирует случайные числа и отправляет их в канал данных.
func writer2(out chan<- any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		out <- rand.Int()
	}
}

// Функция worker2 читает данные из канала данных и выводит их на экран.
func worker2(in <-chan any, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-in:
			if ok {
				fmt.Println(val)
			}
		}
	}
}
