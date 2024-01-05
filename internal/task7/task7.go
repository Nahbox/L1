package task7

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap представляет безопасный для конкурентного доступа map.
type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

// NewSafeMap создает новый экземпляр SafeMap.
func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

// Set устанавливает значение в map.
func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

// Get возвращает значение из map по ключу.
func (sm *SafeMap) Get(key string) int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return sm.data[key]
}

func Task7() {
	// Создаем новый экземпляр SafeMap.
	safeMap := NewSafeMap()

	// Используем WaitGroup для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Запускаем несколько горутин для конкурентной записи данных в map.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			// Генерируем уникальный ключ для каждой горутины.
			key := fmt.Sprintf("key%d", index)

			// Задержка для наглядности.
			time.Sleep(time.Millisecond * 100)

			// Устанавливаем значение в map.
			safeMap.Set(key, index)
			fmt.Printf("goroutine %d wrote to map\n", index)
		}(i)
	}

	// Ожидаем завершения всех горутин.
	wg.Wait()

	// Выводим значения из map после завершения горутин.
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		value := safeMap.Get(key)
		fmt.Printf("Value for key %s: %d\n", key, value)
	}
}
