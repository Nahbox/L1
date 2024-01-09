package task13

import "fmt"

// Функция Task13 демонстрирует два способа обмена значениями переменных.
func Task13() {
	// Исходные значения переменных.
	a := 1
	b := 2

	// Вывод значений до применения первого способа.
	fmt.Printf("До 1-го способа: a=%d, b=%d	", a, b)

	// 1-й способ обмена значениями переменных.
	a, b = b, a

	// Вывод значений после применения первого способа.
	fmt.Printf("После 1-го способа: a=%d, b=%d\n", a, b)

	// Восстановление исходных значений переменных.
	a = 1
	b = 2

	// Вывод значений до применения второго способа.
	fmt.Printf("До 2-го способа: a=%d, b=%d	", a, b)

	// 2-й способ обмена значениями переменных.
	a = a + b
	b = a - b
	a = a - b

	// Вывод значений после применения второго способа.
	fmt.Printf("После 2-го способа: a=%d, b=%d\n", a, b)
}
