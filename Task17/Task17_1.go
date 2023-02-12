package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(val int, array []int) int {
	// если массив пустой, возвращаем - 1
	if len(array) < 1 {
		return -1
	}
	// выставляем границы
	left, right := 0, len(array)-1
	// идем пока индексы не сойдутся
	for left <= right {
		// Берем за пивот центр
		pivot := (left + right) / 2
		//Если искомый больше элемента массива, то ставим пивот на левой границе - сдвигаем впрао
		if array[pivot] < val {
			left = pivot + 1
			//Если искомый меньше элемента массива, то ставим пивот на правой границе - сдвигаем влево
		} else if array[pivot] > val {
			right = pivot - 1
		} else {
			//при совпадении возвращаем
			return pivot
		}
	}
	// не нашли - вовзращаем -1
	return -1
}

func main() {
	rand.Seed(time.Now().Unix())
	array := make([]int, 0)
	//заполняем массив рандомными цифрами
	for i := 0; i < 10; i++ {
		array = append(array, rand.Intn(100))
	}
	//сортируем массив
	sort.Ints(array)
	fmt.Println(array)
	for i := 0; i < 10; i++ {
		if binarySearch(array[i], array) != i {
			//Проверка если вдруг в массиве есть два одинаковых числа
			if (i == 0) && (array[i+1] == array[i]) {
				continue
			} else if (array[i-1] == array[i]) || (array[i+1] == array[i]) {
				continue
			}

			fmt.Println("Search not correct")
			fmt.Println(array[i])
			fmt.Println(i)
			return
		}
	}
	fmt.Println("Search correct")
}
