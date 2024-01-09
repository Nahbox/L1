package task10

import (
	"fmt"
	"sort"
)

// Функция Task10 демонстрирует группировку температурных данных в соответствии с заданным условием.
func Task10() {
	// Исходные температурные данные.
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Группировка температурных данных.
	groups := addToGroup(temperature)

	// Вывод результатов группировки.
	for idx, group := range groups {
		fmt.Println("Группа", idx+1, ":", group)
	}
}

// Функция addToGroup группирует температурные данные в соответствии с заданным условием (разница между температурами <= 10).
func addToGroup(temperature []float64) [][]float64 {
	// Срез для хранения групп температур.
	groups := make([][]float64, 0, 1)
	groupsCount := 0

	// Сортировка температурных данных.
	sort.Float64s(temperature)

	// Группировка температур.
	for _, t := range temperature {
		// Если это первая группа, инициализируем новую группу и добавляем текущую температуру.
		if len(groups) == 0 {
			groups = append(groups, make([]float64, 0))
			groups[groupsCount] = append(groups[groupsCount], t)
			groupsCount++
		} else {
			// Если разница между текущей температурой и последней в текущей группе <= 10,
			// добавляем текущую температуру в текущую группу.
			if t-groups[groupsCount-1][0] <= 10 {
				groups[groupsCount-1] = append(groups[groupsCount-1], t)
			} else {
				// Если разница > 10, создаем новую группу и добавляем текущую температуру в нее.
				groups = append(groups, make([]float64, 0))
				groups[groupsCount] = append(groups[groupsCount], t)
				groupsCount++
			}
		}
	}

	return groups
}
