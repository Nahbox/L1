package task5

import (
	"fmt"
	"math/rand"
	"time"
)

// Каналы отмены и time.Sleep на указанный интервал
func Task5_1() {
	var secCount int
	fmt.Print("Количество секунд: ")
	fmt.Scan(&secCount)
	start := time.Now()

	cancel := make(chan struct{})
	data := make(chan any)

	go writer(data, cancel)
	go worker(data, cancel)

	time.Sleep(time.Duration(secCount) * time.Second)
	cancel <- struct{}{}
	fmt.Println("\nВремя выполнения: ", time.Since(start))
}

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
