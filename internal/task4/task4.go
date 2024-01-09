package task4

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Точка входа в Task4
func Task4() {
	// Переменная для хранения количества воркеров.
	var workersCount int

	var wg sync.WaitGroup
	// Канал для передачи данных между воркерами и писателем.
	data := make(chan any)

	// Создание контекста с функцией отмены.
	ctx, cancel := context.WithCancel(context.Background())

	// Запуск горутины для писателя.
	go func() {
		defer wg.Done()
		writer(ctx, data)
	}()

	// Запуск горутины для обработки сигнала завершения.
	wg.Add(1)
	go func() {
		defer wg.Done()
		checkCtrlC(cancel)
	}()

	// Запуск воркеров в соответствии с указанным количеством.
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			worker(ctx, data, idx)
		}(i + 1)
	}

	// Ожидание завершения всех горутин.
	wg.Wait()
}

// Функция checkCtrlC обрабатывает сигналы завершения (Ctrl+C) и вызывает функцию отмены контекста.
func checkCtrlC(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	cancel()
}

// Функция writer генерирует случайные числа и отправляет их в канал.
func writer(ctx context.Context, out chan<- any) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			out <- rand.Int()
		}
	}
}

// Функция worker обрабатывает данные из канала и выводит их на экран.
func worker(ctx context.Context, in <-chan any, workerNum int) {
	for val := range in {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Воркер №%d, Вывод: ", workerNum)
			fmt.Println(val)
		}
	}
}
