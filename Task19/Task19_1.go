package main

import "fmt"

func main() {
	fmt.Println(backToFrond("kaif"))
	fmt.Println(backToFrond("главрыба"))
	fmt.Println(backToFrond("我厌倦了做这些任务"))
}

//в функции делаем стринг рун и в обратном порядке собираем стрингу
func backToFrond(str string) string {
	var temp string
	runes := []rune(str)
	for i := len(runes) - 1; i > -1; i-- {
		temp = temp + string(runes[i])
	}
	return temp
}
