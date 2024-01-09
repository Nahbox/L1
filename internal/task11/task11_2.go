package task11

import "fmt"

// Функция Task11_2 демонстрирует объединение двух множеств с использованием сортировки слиянием (MergeSort).
func Task11_2() {
	// Исходные множества.
	set1 := []int{1, 7, 4, 12, 3}
	set2 := []int{2, 9, 1, 5, 4, 3, 11, 15}

	// Объединение множеств с помощью сортировки слиянием.
	set1 = append(set1, set2...)
	set2 = nil

	mergeSet := mergeSort(set1)
	fmt.Println(mergeSet)
}

// Функция mergeSort выполняет сортировку слиянием для переданного среза целых чисел.
func mergeSort(items []int) []int {
	// Если длина среза меньше 2, возвращаем его как отсортированный.
	if len(items) < 2 {
		return items
	}

	// Рекурсивно сортируем левую и правую половины среза.
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])

	// Объединяем отсортированные половины.
	return merge(first, second)
}

// Функция merge объединяет два отсортированных среза в один отсортированный.
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0

	// Сравниваем элементы и добавляем минимальный в результирующий срез.
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	// Добавляем оставшиеся элементы из левой половины (если есть).
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	// Добавляем оставшиеся элементы из правой половины (если есть).
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
