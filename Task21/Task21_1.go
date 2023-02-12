package main

import (
	"fmt"
	"strings"
)

func main() {
	//создаем экземпляр интерфейса Runes которому удовлетворяет структура Adapter
	var run Runes = &Adapter{}
	//вызываем метод
	temp := run.GetHugeRune(12)
	temp[3] = 'a'
	fmt.Println(string(temp))
}

//у нас есть сторонний пакет HugeString с методом GetHugeString

type HugeString struct {
}

func (s *HugeString) GetHugeString(len int) string {
	return strings.Repeat("у", len)
}

//Но в нашем сервисе мы используем руны

type Runes interface {
	GetHugeRune(len int) []rune
}

//и мы создаем адаптер, который имеет экземпляр структуры из стороннего пакета

type Adapter struct {
	str HugeString
}

//и функцию с помощью которой адаптируем

func (a *Adapter) GetHugeRune(len int) []rune {
	return []rune(a.str.GetHugeString(len))
}
