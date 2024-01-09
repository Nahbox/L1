package task22

import (
	"fmt"
	"math/big"
)

// Функция Task22 реализует операции умножения, деления, сложения и вычитания с использованием пакета math/big.
func Task22() {
	// Инициализация переменных типа big.Int
	var a, b big.Int

	// Ввод значения переменной a
	fmt.Print("Введите значение a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Ввод значения переменной b
	fmt.Print("Введите значение b: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	// Умножение
	multiplicationResult := new(big.Int).Mul(&a, &b)
	fmt.Printf("%s * %s = %s\n", &a, &b, multiplicationResult)

	// Деление
	if b.Sign() != 0 {
		divisionResult := new(big.Int).Quo(&a, &b)
		fmt.Printf("%s / %s = %s\n", &a, &b, divisionResult)
	} else {
		fmt.Println("Нельзя делить на ноль")
	}

	// Сложение
	additionResult := new(big.Int).Add(&a, &b)
	fmt.Printf("%s + %s = %s\n", &a, &b, additionResult)

	// Вычитание
	subtractionResult := new(big.Int).Sub(&a, &b)
	fmt.Printf("%s - %s = %s\n", &a, &b, subtractionResult)
}
