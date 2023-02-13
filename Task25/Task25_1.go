package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Second())
	sleep(time.Second * 2)
	fmt.Println(time.Now().Second())
}

//используем метод After, в котором через определнное время поступает сигнал из канала
func sleep(d time.Duration) {
	<-time.After(d)
}
