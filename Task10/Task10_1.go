package main

import "fmt"

func main() {
	//входные данные
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	//делаем мапу
	m := make(map[int][]float64)

	for _, value := range input {
		//избавляемся от дробных и единиц
		key := (int(value) / 10) * 10
		//добавляем в группу
		m[key] = append(m[key], value)
	}
	//печатаем
	fmt.Println(m)
}
