package main

import "fmt"

func main() {
	a := 2023
	b := 111
	fmt.Println("a = ", a, " b =", b)
	//я кстати не знаю сколько это действие памяти занимает...
	a, b = b, a
	fmt.Println("a = ", a, " b =", b)
}
