package main

import "fmt"

func main() {
	var number int64 = 123131212312
	var mask int64
	//выбираем на какое число будем сдвигать единицу
	i := 4
	// маска с единицей
	mask = (1 << i)
	//выставляем 1
	number |= mask
	//выводим в бинарном виде
	fmt.Printf("%b \n", number)

	//выставляем 0
	number &= ^(mask)
	//выводим в бинарном виде
	fmt.Printf("%b", number)

}
