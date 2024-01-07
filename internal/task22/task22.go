package task22

import (
	"fmt"
	"math/big"
)

// Реализация с использованием пакета math/big
func Task22() {
	var a, b big.Int

	fmt.Print("Введите значение a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

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
		divisionResult := new(big.Int).Div(&a, &b)
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
