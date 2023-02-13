package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// устанавливаем время завершения в секундах
	n := 10

	//создаем контекст с заврешением по времени
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(n*int(time.Second)))

	// создаем канал с данными
	ch := make(chan int)

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//воркер который принимает данные с каналов
	go func() {
		//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
		defer wg.Done()
		//бесконечный цикл, ожидаем прием данных или заверщение по контексту
		for {
			select {
			case a := <-ch:
				fmt.Printf("Received : %v \n", a)
			case <-ctx.Done():
				fmt.Println("Receiver shutdown")
				return
			}
		}
	}()

	//ворекер который слушает завершение контекста и отправляет данные в канал
	go func() {
		i := 0
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Sender shutdown")
				return
			default:
				ch <- i
				i++
			}
			time.Sleep(time.Second)
		}
	}()
	//ждем завершения воркеров
	wg.Wait()
	log.Println("Process shutdown")
}
