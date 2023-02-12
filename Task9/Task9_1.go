package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//контекст для завершения
	ctx, cancel := context.WithCancel(context.Background())

	//два канала для конвеера
	ch1 := make(chan int)
	ch2 := make(chan int)

	//горутина с воркером, который принимает число с 1 канала и отправляет,умножив на 2, во 2 канал
	go func() {
		for {
			select {
			case ch2 <- (<-ch1 * 2):

			case <-ctx.Done():
				return
			}
		}
	}()
	//горутинка которая принимает число из 2 канала и выводит его
	go func() {
		for {
			select {
			case a := <-ch2:
				fmt.Println(a)
			case <-ctx.Done():
				return
			}
		}
	}()
	//пишем в канал 1
	for i := 0; i < 4; i++ {
		ch1 <- rand.Intn(100)
		time.Sleep(time.Second)
	}
	//завершаем горутинки
	cancel()

}
