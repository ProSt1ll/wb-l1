package main

import "fmt"

func main() {
	//входные данные
	num1 := []int{7, 2, 3, 4, 5}
	num2 := []int{2, 3, 6, 8}
	str1 := []string{"nil", "cat", "dog", "test", "wb"}
	str2 := []string{"nul", "cat", "dog", "net", "wb"}
	//вывод
	fmt.Println(intersection(num1, num2))
	fmt.Println(intersection(str1, str2))

}

//всем привет и сегодня мы будем использовать джинерики
func intersection[T any](input1, input2 []T) (output []T) {
	//делаем мапу с интерфейсным ключом и значением пустой структуры
	m := make(map[interface{}]struct{})
	//бежим по слайсу и пишем в мапу
	for _, value := range input1 {
		m[value] = struct{}{}
	}
	//бежим по второму массиву и если находим значение в мапе то добавляем в возвращаемый массив
	for _, value := range input2 {
		_, ok := m[value]
		if ok {
			output = append(output, value)
		}
	}
	return output
}
