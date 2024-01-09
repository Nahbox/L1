package task12

import "fmt"

// Функция Task12 демонстрирует использование структуры Set.
func Task12() {
	// Создание нового множества.
	set := NewSet()

	// Добавление элементов в множество.
	set.Add("cat", "cat", "dog", "cat", "tree")

	// Удаление элемента из множества.
	set.Delete("tree")

	// Проверка наличия элемента в множестве.
	fmt.Println(set.Contain("dog"))

	// Получение всех ключей множества.
	fmt.Println(set.GetKeys())

	// Вывод содержимого множества.
	fmt.Println(set)
}

// Тип Set представляет собой простую реализацию множества.
type Set struct {
	data map[any]int
}

// NewSet создает новый экземпляр множества.
func NewSet() *Set {
	return &Set{data: map[any]int{}}
}

// Add добавляет элементы в множество.
func (s *Set) Add(values ...any) {
	for _, value := range values {
		s.data[value]++
	}
}

// Delete удаляет элементы из множества.
func (s *Set) Delete(keys ...any) {
	for _, key := range keys {
		if s.data[key] > 1 {
			s.data[key]--
		} else {
			delete(s.data, key)
		}
	}
}

// Contain проверяет наличие элемента в множестве.
func (s Set) Contain(data any) bool {
	_, ok := s.data[data]
	return ok
}

// GetKeys возвращает все ключи множества.
func (s *Set) GetKeys() []any {
	keys := make([]any, 0, len(s.data))

	for k, v := range s.data {
		if v > 0 {
			keys = append(keys, k)
		}
	}

	return keys
}
