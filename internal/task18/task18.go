package task18

import (
	"fmt"
	"sync"
)

func Task18() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Используем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Запускаем 100 горутин для инкрементирования счетчика
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
			fmt.Println("Counter Value:", counter.GetValue())
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}

// Counter - структура счетчика
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment - инкрементирует значение счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue - возвращает текущее значение счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}
