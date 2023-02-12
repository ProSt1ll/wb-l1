package main

import "fmt"

//родительская структура
type Human struct {
	name string
}

// дочерняя структура
type Action struct {
	Human
}

// метод родительской структуры
func (h *Human) Talk() {
	fmt.Println("Hi, my name is " + h.name)
}

// Action has all Human methods
func main() {
	//инициализация дочерней структуры
	child := Action{
		Human: Human{
			name: "John",
		},
	}
	//вызов метода родительской структуры с помощью дочерней
	child.Talk()
}
