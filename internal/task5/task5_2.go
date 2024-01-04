package task5

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// С использованием context.WithTimeout
func Task5_2() {
	var secCount int
	fmt.Print("Количество секунд: ")
	fmt.Scan(&secCount)
	start := time.Now()

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(secCount)*time.Second)
	data := make(chan any)

	go writer2(data, ctx)
	go worker2(data, ctx)

	time.Sleep(time.Duration(secCount) * time.Second)
	fmt.Println("\nВремя выполнения: ", time.Since(start))
}

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
