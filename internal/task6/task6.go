package task6

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Task6() {
	var wg sync.WaitGroup

	// Канал отмены
	cancelCh := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		forCancelCh(cancelCh)
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Завершение горутины при помощи канала отмены")
	cancelCh <- struct{}{}

	// Контекст
	ctxWCancel, cancel := context.WithCancel(context.Background()) // WithCancel

	wg.Add(1)
	go func() {
		defer wg.Done()
		forCtx(ctxWCancel)
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("Завершение горутины с использованием контекста контекста (WithCancel)")
	cancel()

	ctxWTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second) // WithTimeout

	wg.Add(1)
	go func() {
		defer wg.Done()
		forCtx(ctxWTimeout)
	}()

	wg.Wait()

	fmt.Println("Завершение горутины с использованием контекста (WithTimeout)")
}

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
