package task11

import "fmt"

// Объединение с помощью MergeSort
func Task11_2() {
	set1 := []int{1, 7, 4, 12, 3}
	set2 := []int{2, 9, 1, 5, 4, 3, 11, 15}

	set1 = append(set1, set2...)
	set2 = nil

	mergeSet := mergeSort(set1)
	fmt.Println(mergeSet)
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
