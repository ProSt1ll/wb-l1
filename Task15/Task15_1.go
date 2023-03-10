package main

import (
	"fmt"
	"strings"
)

//В данном случае возможна ошибка на строке 23, а именно последний символ может быть некорректен
//тк string это по сути срез байт, который недоступен для изменения, с кодировкой UTF-8(каждый символ занимает 3 байта)
// и если как в данном случае выделить интервал, то последний символ обрежется на 1 байте и он будет невалидным
//на примере как раз видим что у нас получилось 33 корректных символа - 99 байт + 1 некорректный
//Также можно отметить что интервал не соответствует количеству символов, что тоже стоит учесть
//не вижу смысла контролировать начало интервала и конец кратности 3, легче перевести строку в срез рун и все будет
//отлично :)

var justString string

func someFunc() {
	//1 << 10 обычное число инт, просто единицу сдвинули на 10 знаков
	str := createHugeString(1 << 10)
	//берем как в примере и получаем ошибку последнего символа
	justString = str[:100]
	fmt.Println(justString)
	//берем интервал уже от рун и все хорошо
	justString = string([]rune(str)[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}

func createHugeString(len int) string {
	return strings.Repeat("字", len)
}
