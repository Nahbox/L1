package task9

import (
	"fmt"
)

// Функция Task9 демонстрирует использование каналов для передачи данных между горутинами.
func Task9() {
	// Создание двух каналов для передачи данных между горутинами.
	ch1 := make(chan int)
	ch2 := make(chan int)
	// Срез чисел для передачи в первый канал.
	nums := []int{2, 4, 6, 8, 10}

	// Запуск горутины для записи чисел в первый канал.
	go func() {
		writerX(ch1, nums)
	}()

	// Запуск горутины для возведения чисел в квадрат и передачи результата во второй канал.
	go func() {
		powX(ch1, ch2)
	}()

	// Вывод результатов из второго канала.
	for xpow := range ch2 {
		fmt.Println(xpow)
	}
}

// Функция writerX записывает числа из среза в канал.
func writerX(xchan chan<- int, nums []int) {
	defer close(xchan)
	for _, val := range nums {
		xchan <- val
	}
}

// Функция powX возводит числа в квадрат и передает результат во второй канал.
func powX(xchan <-chan int, xpowchan chan<- int) {
	defer close(xpowchan)
	for x := range xchan {
		xpowchan <- x * x
	}
}
