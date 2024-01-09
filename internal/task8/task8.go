package task8

import (
	"fmt"
)

// Функция Task8 позволяет пользователю устанавливать бит в 1 или 0 на указанной позиции в числе типа int64.
func Task8() {
	// Переменные для ввода данных.
	var num int64
	var position uint
	var value bool

	// Ввод данных.
	fmt.Print("Введите число (int64): ")
	fmt.Scan(&num)

	fmt.Print("Введите позицию бита: ")
	fmt.Scan(&position)

	fmt.Print("Установить бит в 1 или 0 (true/false): ")
	fmt.Scan(&value)

	var result int64
	// Вызываем соответствующую функцию в зависимости от значения value.
	if value {
		result = setBitOne(num, position)
	} else {
		result = setBitZero(num, position)
	}

	// Выводим результат.
	fmt.Printf("Результат: %d\n", result)
}

// Функция setBitOne устанавливает бит в 1 на указанной позиции.
func setBitOne(num int64, position uint) int64 {
	return num | (1 << position)
}

// Функция setBitZero устанавливает бит в 0 на указанной позиции.
func setBitZero(num int64, position uint) int64 {
	return num & ^(1 << position)
}
