package task12

import "fmt"

func Task12() {
	set := NewSet()

	set.Add("cat", "cat", "dog", "cat", "tree")
	set.Delete("tree")

	fmt.Println(set.Contain("dog"))
	fmt.Println(set.GetKeys())
	fmt.Println(set)
}

type Set struct {
	data map[any]int
}

func NewSet() *Set {
	return &Set{data: map[any]int{}}
}

func (s *Set) Add(values ...any) {
	for _, value := range values {
		s.data[value]++
	}
}

func (s *Set) Delete(keys ...any) {
	for _, key := range keys {
		if s.data[key] > 1 {
			s.data[key]--
		} else {
			delete(s.data, key)
		}
	}
}

func (s Set) Contain(data any) bool {
	_, ok := s.data[data]
	return ok
}

func (s *Set) GetKeys() []any {
	keys := make([]any, 0, len(s.data))

	for k, v := range s.data {
		if v > 0 {
			keys = append(keys, k)
		}
	}

	return keys
}
