package task16

import "fmt"

// Функция Task16 демонстрирует использование алгоритма быстрой сортировки (QuickSort).
func Task16() {
	// Исходный набор данных.
	set := []int{1, 7, 4, 12, 3, 2, 9, 1, 5, 4, 3, 11, 15}

	// Вызов функции для сортировки массива.
	quickSortStart(set)

	// Вывод отсортированного массива.
	fmt.Println(set)
}

// Функция partition выполняет разделение массива для алгоритма QuickSort.
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

// Функция quickSort выполняет рекурсивную сортировку массива методом QuickSort.
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

// Функция quickSortStart предоставляет интерфейс для вызова алгоритма QuickSort с использованием массива.
func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
