package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	//создаем контекст с заврешением по времени
	//вообще мы можем использовать любой context с отменой, смысл от этого не поменяется
	ctx, _ := context.WithTimeout(context.Background(), time.Second*4)

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	//горутина работяги который работает и слушает конекст выполнения
	go func() {
		//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Gorutine shutdown")
				return
			default:
				fmt.Println("work")
			}
			time.Sleep(time.Second)
		}
	}()

	//ждем завершения горутинки
	wg.Wait()

}
