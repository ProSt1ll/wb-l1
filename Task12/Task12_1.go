package main

import "fmt"

func main() {
	//входные данные
	input := []string{"cat", "cat", "dog", "cat", "tree"}

	//создаем мапу
	set := make(map[string]struct{})
	//бежим по входным и добавляем в мапу, нам все равно, если заменит тк это не играет роли...
	//а доставать проверять время занимает
	for _, value := range input {
		set[value] = struct{}{}
	}
	//вывод
	for key, _ := range set {
		fmt.Print(key + " ")
	}
}
