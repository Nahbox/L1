package task11

import (
	"fmt"
	"math"
)

// Функция Task11_1 демонстрирует объединение двух множеств с использованием подсчета.
func Task11_1() {
	// Исходные множества.
	set1 := []int{1, 7, 4, 12, 3}
	set2 := []int{2, 9, 1, 5, 4, 3, 11, 15}

	// Получение результатов подсчета и минимального значения из множеств.
	countSet, minVal := counter(set1, set2)

	// Объединение множеств и вывод результата.
	result := mergeSets(countSet, minVal, len(set1)+len(set2))
	fmt.Println(result)
}

// Функция counter выполняет подсчет элементов в каждом из множеств и возвращает результаты подсчета и минимальное значение.
func counter(set1, set2 []int) ([]int, int) {
	// Инициализация переменных для максимального и минимального значений.
	maxVal := math.MinInt
	minVal := math.MaxInt

	// Находим максимальное и минимальное значения в обоих множествах.
	for _, val := range set1 {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}
	for _, val := range set2 {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}
	}

	// Инициализация счетчика для каждого элемента в интервале от минимального до максимального значения.
	countSet := make([]int, maxVal-minVal+1)

	// Подсчет элементов в первом множестве.
	for _, val := range set1 {
		countSet[val-minVal]++
	}

	// Подсчет элементов во втором множестве.
	for _, val := range set2 {
		countSet[val-minVal]++
	}

	return countSet, minVal
}

// Функция mergeSets объединяет множества на основе результатов подсчета и минимального значения.
func mergeSets(countSet []int, minVal int, totalLen int) []int {
	// Инициализация результирующего множества.
	resSet := make([]int, 0, totalLen)

	// Проход по каждому элементу в результате подсчета.
	for idx, val := range countSet {
		// Добавление элемента в результирующее множество столько раз, сколько он встречается в подсчете.
		for i := 0; i < val; i++ {
			resSet = append(resSet, idx+minVal)
		}
	}

	return resSet
}
