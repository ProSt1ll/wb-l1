package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	// устанавливаем время завершения
	n := 10

	//создаем канал с заврешением
	quit := make(chan struct{})

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
			case a, ok := <-ch:
				if !ok {
					fmt.Println("Receiver shutdown")
					return
				}
				fmt.Printf("Received : %v \n", a)
			}
		}
	}()

	//воркер который слушает канал завершения и отправляет данные в канал
	go func() {
		i := 0
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Sender shutdown")
				close(ch)
				return
			default:
				ch <- i
				i++
			}
			time.Sleep(time.Second)
		}
	}()

	// ждем n секунд
	time.Sleep(time.Duration(n * int(time.Second)))
	//отправляем в канал завершения
	quit <- struct{}{}
	//ждем завершение воркеров
	wg.Wait()
	log.Println("Process shutdown")
}
