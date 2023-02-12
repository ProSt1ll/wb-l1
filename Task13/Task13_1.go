package main

import "fmt"

func main() {
	a := 2023
	b := 1111
	fmt.Println("a =", a, "b =", b)
	//ловкость рук никакой магии
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("a =", a, "b =", b)
}
