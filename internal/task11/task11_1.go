package task11

import (
	"fmt"
	"math"
)

// Объединение подсчетом
func Task11_1() {
	set1 := []int{1, 7, 4, 12, 3}
	set2 := []int{2, 9, 1, 5, 4, 3, 11, 15}

	countSet, minVal := counter(set1, set2)
	result := mergeSets(countSet, minVal, len(set1)+len(set2))

	fmt.Println(result)
}

func counter(set1, set2 []int) ([]int, int) {
	maxVal := math.MinInt
	minVal := math.MaxInt
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

	countSet := make([]int, maxVal-minVal+1)

	for _, val := range set1 {
		countSet[val-minVal]++
	}
	for _, val := range set2 {
		countSet[val-minVal]++
	}

	return countSet, minVal
}

func mergeSets(countSet []int, minVal int, totalLen int) []int {
	resSet := make([]int, 0, totalLen)

	for idx, val := range countSet {
		for i := 0; i < val; i++ {
			resSet = append(resSet, idx+minVal)
		}
	}
	return resSet
}
