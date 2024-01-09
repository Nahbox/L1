package task23

import (
	"errors"
	"fmt"
	"log"
)

func Task23() {
	mySlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Original Slice:", mySlice)

	// Удаление элемента по индексу
	mySlice, err := removeElement(mySlice, 3)
	if err != nil {
		log.Print(err)
	}

	fmt.Println("Slice after removal:", mySlice)
}

// removeElement удаляет элемент из слайса по указанному индексу.
// Если индекс некорректен, возвращает оригинальный слайс и ошибку.
func removeElement(slice []int, index int) ([]int, error) {
	// Проверка на корректность индекса
	if index < 0 || index >= len(slice) {
		return slice, errors.New("Invalid index")
	}

	// Перераспределение слайса
	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]

	return slice, nil
}
