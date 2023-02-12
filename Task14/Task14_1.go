package main

import (
	"fmt"
	"reflect"
)

func main() {
	//вывод
	fmt.Println(Type("kaif"))
	fmt.Println(Type(123))
	fmt.Println(Type(123.3))
}

//используем пакет reflect

func Type(val interface{}) reflect.Type {
	return reflect.TypeOf(val)
}
