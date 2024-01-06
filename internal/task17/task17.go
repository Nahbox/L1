package task17

import "fmt"

func Task17() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 7
	result := binarySearch(arr, target)

	if result == -1 {
		fmt.Println("target не найден")
	} else {
		fmt.Printf("target найден: index=%d\n", result)
	}
}

func binarySearch(arr []int, target int) int {
	// Нижняя и верхняя границы
	low, high := 0, len(arr)-1

	for low <= high { // Продолжать поиск до тех пор, пока low <= high
		// Рассчитать средний индекс
		mid := low + (high-low)/2
		// Проверка, найден ли элемент
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			// Если target больше, сосредоточьтесь на правой половине
			low = mid + 1
		} else {
			// Если target меньше, сосредоточьтесь на левой половине
			high = mid - 1
		}
	}

	// Если элемент не найден, return -1
	return -1
}
