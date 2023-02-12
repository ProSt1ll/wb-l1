package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//создаем канал с заврешением
	quit := make(chan struct{})

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	//горутина работяги который работает и слушает канал выхода
	go func() {
		//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Gorutine shutdown")
				return
			default:
				fmt.Println("work")
			}
			time.Sleep(time.Second)
		}
	}()
	//Немного ждем
	time.Sleep(time.Second * 3)
	//отправляем пустую структуру в канал завершения
	quit <- struct{}{}
	//ждем завершения горутинки
	wg.Wait()

}
