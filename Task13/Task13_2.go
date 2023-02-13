package main

import "fmt"

func main() {
	a := 2023
	b := 111
	fmt.Println("a = ", a, " b =", b)
	//свап по Goлогновски
	a, b = b, a
	fmt.Println("a = ", a, " b =", b)
}
