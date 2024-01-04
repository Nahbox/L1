package task1

import "fmt"

type Action struct {
	Human
}

type Human struct {
	age  int
	name string
}

func (h Human) sayName() {
	fmt.Println("My name is", h.name)
}

func (h Human) powAge() {
	fmt.Println(h.age * h.age)
}

func Task1() {
	human := Human{age: 20, name: "John"}
	action := Action{Human: human}
	action.powAge()
	action.sayName()
}
