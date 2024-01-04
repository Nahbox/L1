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

func Task4_1() {
	var workersCount int
	var wg sync.WaitGroup
	data := make(chan any)

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Print("Количество воркеров: ")
	fmt.Scan(&workersCount)

	go func() {
		defer wg.Done()
		writer(ctx, data)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		checkCtrlC(cancel)
	}()

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			worker(ctx, data, idx)
		}(i + 1)
	}

	wg.Wait()
}

func checkCtrlC(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	cancel()
}

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
