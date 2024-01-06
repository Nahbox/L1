package task25

import (
	"fmt"
	"time"
)

func Task25() {
	fmt.Println("Start")
	sleep(2000 * time.Millisecond) // Приостановка выполнения на 2 секунды (2000 миллисекунд)
	fmt.Println("End")
}

// sleep реализует функцию задержки с использованием каналов и select
func sleep(d time.Duration) {
	ch := make(chan bool)

	go func() {
		// Запускаем таймер и ждем указанное количество миллисекунд
		timer := time.NewTimer(d)
		defer timer.Stop()

		// Ожидаем события от таймера или закрытия канала
		select {
		case <-timer.C:
		case <-ch:
		}
		close(ch)
	}()

	// Ждем, пока канал не будет закрыт
	<-ch
}
