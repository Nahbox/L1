package task9

import (
	"fmt"
)

func Task9() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{2, 4, 6, 8, 10}

	go func() {
		writerX(ch1, nums)
	}()

	go func() {
		powX(ch1, ch2)
	}()

	for xpow := range ch2 {
		fmt.Println(xpow)
	}
}

func writerX(xchan chan<- int, nums []int) {
	for _, val := range nums {
		xchan <- val
	}
	close(xchan)
}

func powX(xchan <-chan int, xpowchan chan<- int) {
	for x := range xchan {
		xpowchan <- x * x
	}
	close(xpowchan)
}
