package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkStr("name"))
	fmt.Println(checkStr("LAlko"))
}

//функция преобразует в срез рун и проходится по нему,
//если вытаскивает из мапы - повтор
func checkStr(str string) bool {
	m := make(map[rune]struct{})
	str = strings.ToLower(str)
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		_, ok := m[runes[i]]
		if ok {
			return false
		} else {
			m[runes[i]] = struct{}{}
		}
	}
	return true
}
