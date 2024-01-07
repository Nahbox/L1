package task8

import (
	"fmt"
)

func Task8() {
	var num int64
	var position uint
	var value bool

	// Ввод данных
	fmt.Print("Введите число (int64): ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита: ")
	fmt.Scan(&position)

	fmt.Print("Установить бит в 1 или 0 (true/false): ")
	fmt.Scan(&value)

	var result int64
	if value {
		result = setBitOne(num, position)
	} else {
		result = setBitZero(num, position)
	}

	fmt.Printf("Результат: %d\n", result)
}

func setBitOne(num int64, position uint) int64 {
	return num | (1 << position)
}

func setBitZero(num int64, position uint) int64 {
	return num & ^(1 << position)
}
