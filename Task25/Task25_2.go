package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Second())
	sleeps(time.Second * 2)
	fmt.Println(time.Now().Second())
}

//ждем в цикле пока не настанет момент: время вызова функции + заданное время
func sleeps(d time.Duration) {
	now := time.Now().Unix()
	for time.Now().Unix()-now < int64(d.Seconds()) {
	}
}
