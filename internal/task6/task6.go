package task6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция Task6 демонстрирует использование канала отмены и контекста для завершения горутин.
func Task6() {
	var wg sync.WaitGroup

	// Канал отмены.
	cancelCh := make(chan struct{})

	// Запуск горутины для использования канала отмены.
	wg.Add(1)
	go func() {
		defer wg.Done()
		forCancelCh(cancelCh)
	}()

	// Ожидание 5 секунд и отправка сигнала отмены в канал.
	time.Sleep(5 * time.Second)
	fmt.Println("Завершение горутины при помощи канала отмены")
	cancelCh <- struct{}{}

	// Контекст с функцией отмены (WithCancel).
	ctxWCancel, cancel := context.WithCancel(context.Background())

	// Запуск горутины для использования контекста с функцией отмены (WithCancel).
	wg.Add(1)
	go func() {
		defer wg.Done()
		forCtx(ctxWCancel)
	}()

	// Ожидание 5 секунд и вызов функции отмены контекста.
	time.Sleep(5 * time.Second)
	fmt.Println("Завершение горутины с использованием контекста (WithCancel)")
	cancel()

	// Контекст с таймаутом (WithTimeout).
	ctxWTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Запуск горутины для использования контекста с таймаутом (WithTimeout).
	wg.Add(1)
	go func() {
		defer wg.Done()
		forCtx(ctxWTimeout)
	}()

	// Ожидание завершения всех горутин.
	wg.Wait()

	fmt.Println("Завершение горутины с использованием контекста (WithTimeout)")
}

// Функция forCancelCh ожидает сигнала отмены через канал отмены.
func forCancelCh(cancelCh <-chan struct{}) {
	fmt.Println("Горутина запущена...")
	for {
		select {
		case <-cancelCh:
			return
		default:
		}
	}
}

// Функция forCtx ожидает завершения контекста.
func forCtx(ctx context.Context) {
	fmt.Println("Горутина запущена...")
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
	}
}
