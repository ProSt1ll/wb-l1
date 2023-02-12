package main

import (
	"fmt"
)

func main() {
	//вывод
	fmt.Println(PrinType("kaif"))
	fmt.Println(PrinType(123))
	fmt.Println(PrinType(123.3))
}

//используем Sprintf

func PrinType(val interface{}) string {
	return fmt.Sprintf("%T", val)
}
