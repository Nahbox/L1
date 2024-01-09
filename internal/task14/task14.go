package task14

import "fmt"

// Функция Task14 демонстрирует использование интерфейса для определения типа переменной.
func Task14() {
	var integer int
	var str string
	var boolean bool
	var chInt chan int
	var chString chan string
	var chBool chan bool

	// Вызов функции для определения типов переменных.
	calculateType(chBool)   // chan bool
	calculateType(str)      // string
	calculateType(boolean)  // bool
	calculateType(chString) // chan string
	calculateType(chInt)    // chan int
	calculateType(integer)  // int
}

// Функция calculateType определяет тип переменной и выводит соответствующее сообщение.
func calculateType(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("Переменная типа int")
	case string:
		fmt.Println("Переменная типа string")
	case bool:
		fmt.Println("Переменная типа bool")
	case chan int:
		fmt.Println("Переменная типа chan int")
	case chan string:
		fmt.Println("Переменная типа chan string")
	case chan bool:
		fmt.Println("Переменная типа chan bool")
	}
}
