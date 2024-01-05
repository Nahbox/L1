package task10

import (
	"fmt"
	"sort"
)

func Task10() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := addToGroup(temperature)
	for idx, group := range groups {
		fmt.Println("Группа", idx+1, ":", group)
	}
}

func addToGroup(temperature []float64) [][]float64 {
	groups := make([][]float64, 0, 1)
	groupsCount := 0
	sort.Float64s(temperature)
	for _, t := range temperature {
		if len(groups) == 0 {
			groups = append(groups, make([]float64, 0))
			groups[groupsCount] = append(groups[groupsCount], t)
			groupsCount++
		} else {
			if t-groups[groupsCount-1][0] <= 10 {
				groups[groupsCount-1] = append(groups[groupsCount-1], t)
			} else {
				groups = append(groups, make([]float64, 0))
				groups[groupsCount] = append(groups[groupsCount], t)
				groupsCount++
			}
		}
	}

	return groups
}
