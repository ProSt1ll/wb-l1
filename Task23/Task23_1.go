package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := make([]int, 0)
	//заполняем массив рандомными цифрами
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(101))
	}
	//печатаем исходный срез
	fmt.Println(array)
	//печатаем срез с удаленным элементом
	fmt.Println(deleteElem(array, 3))
}

//функция пробегается по элементам и сдвигает на 1 назад, после чего переопределяет срез на полученный,
//но меньшей длинны
func deleteElem[T any](array []T, k int) []T {
	for i := k; i < len(array)-1; i++ {
		array[i] = array[i+1]
	}
	array = array[:len(array)-1]
	return array
}
