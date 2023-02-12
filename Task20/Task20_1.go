package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(backToFrond("snow dog sun"))
	fmt.Println(backToFrond("глав рыба"))
	fmt.Println(backToFrond("我厌倦了 做这些任务"))
}

//сплитим слова с помощью пакета strings и в обратном порядке склеиваем стрингу
func backToFrond(str string) string {
	var temp string
	words := strings.Split(str, " ")
	for i := len(words) - 1; i > -1; i-- {
		temp += words[i] + " "
	}
	return temp
}
