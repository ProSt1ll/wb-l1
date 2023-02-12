package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func partition(array []int, left, right int) int {
	// берем центральный элемент за пивот
	pivot := array[(left+right)/2]

	idxLeft, idxRight := left, right

	for idxLeft <= idxRight {
		//ищем последний индекс элемента слева , больше чем пивот
		for array[idxLeft] < pivot {
			idxLeft++
		}
		//превый индекс элемента справва, меньше чем пивот
		for array[idxRight] > pivot {
			idxRight--
		}
		//выполняем пока индекс слева не встретиться индексом справа
		if idxLeft >= idxRight {
			break
		}
		//меняем местами, идем дальше
		array[idxLeft], array[idxRight] = array[idxRight], array[idxLeft]
		idxLeft++
		idxRight--
	}
	return idxRight
}

//наша функция сортировки принимает неостортированный массив и границы сортировки
func quickSort(array []int, left, right int) {
	//если массив меньше 2 элементов - сортировать не надо
	if len(array) < 2 {
		return
	}
	//ждем пока наши границы не сойдутся
	if left < right {
		// меняем местами элементы двух частей и получаем новую границу
		idxPart := partition(array, left, right)
		// проходим по левой части разделенного массива
		quickSort(array, left, idxPart)
		// проходим по правой части разделенного массива
		quickSort(array, idxPart+1, right)
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	array := make([]int, 0)
	//заполняем массив рандомными цифрами
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(100))
	}
	fmt.Println(array)
	//сорируем копию массива для проверки
	sorted := array
	sort.Ints(sorted)
	//сортируем
	quickSort(array, 0, len(array)-1)
	fmt.Println(sorted)
	fmt.Println(array)
}
