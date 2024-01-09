package task1

import "fmt"

type Action struct {
	Human
}

type Human struct {
	age  int
	name string
}

// Метод sayName выводит имя человека.
func (h Human) sayName() {
	fmt.Println("My name is", h.name)
}

// Метод powAge выводит квадрат возраста человека.
func (h Human) powAge() {
	fmt.Println(h.age * h.age)
}

// Функция Task1 для выполнения задачи.
func Task1() {
	// Создание объекта типа Human с возрастом 20 и именем "John".
	human := Human{age: 20, name: "John"}
	// Создание объекта типа Action, включающего в себя объект Human.
	action := Action{Human: human}

	// Вызов метода powAge для вывода квадрата возраста.
	action.powAge()
	// Вызов метода sayName для вывода имени.
	action.sayName()
}
