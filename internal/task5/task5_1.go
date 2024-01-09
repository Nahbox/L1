package task5

import (
	"fmt"
	"math/rand"
	"time"
)

// Тип any для использования в каналах.
type any interface{}

// Функция Task5_1 выполняет задачу по использованию каналов отмены и time.Sleep на указанный интервал.
func Task5_1() {
	// Переменная для хранения количества секунд.
	var secCount int
	fmt.Print("Количество секунд: ")
	fmt.Scan(&secCount)
	// Запись текущего времени.
	start := time.Now()

	// Создание каналов отмены и передачи данных.
	cancel := make(chan struct{})
	data := make(chan any)

	// Запуск горутины для писателя.
	go writer(data, cancel)
	// Запуск горутины для воркера.
	go worker(data, cancel)

	// Ожидание указанного интервала времени.
	time.Sleep(time.Duration(secCount) * time.Second)
	// Отправка сигнала отмены в канал отмены.
	cancel <- struct{}{}

	// Вывод времени выполнения.
	fmt.Println("\nВремя выполнения: ", time.Since(start))
}

// Функция writer генерирует случайные числа и отправляет их в канал данных.
func writer(out chan<- any, cancel <-chan struct{}) {
	for {
		select {
		case <-cancel:
			return
		default:
		}
		out <- rand.Int()
	}
}

// Функция worker читает данные из канала данных и выводит их на экран.
func worker(in <-chan any, cancel <-chan struct{}) {
	for {
		select {
		case <-cancel:
			return
		case val, ok := <-in:
			if ok {
				fmt.Println(val)
			}
		}
	}
}
