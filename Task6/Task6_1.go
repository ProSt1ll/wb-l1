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
	wg.Add(2)
	//горутина работяги который работает и слушает канал выхода
	for i := 0; i < 2; i++ {
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
	}
	//Немного ждем
	time.Sleep(time.Second * 3)
	//закрываем канал тк у нас несколько горутин, но стоит учесть
	//что никто не должен в него писать иначе будет паника
	close(quit)

	//ждем завершения горутинки
	wg.Wait()

}
