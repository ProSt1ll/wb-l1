package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// выбираем количество воркеров
	workers := 10

	// канал для завершения
	quit := make(chan os.Signal, 1)

	// канал для завершения горутин
	quitGorutine := make(chan struct{}, 1)

	// связываем сигналы ОС и канала
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// создаем канал с данными
	ch := make(chan int)

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}

	// создаем воркеров в цикле
	for i := 0; i < workers; i++ {

		wg.Add(1)
		// запускаем горутинку воркера
		go func(id int) {
			log.Println("Start gorutine: ", id)
			//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
			defer wg.Done()
			for {
				select {
				// читаем канал завершения горут
				case <-quitGorutine:
					log.Println("Shutdown goroutine: ", id)
					return
				case data := <-ch:
					fmt.Printf("Worker, %v data: %v \n", id, data)
				}
			}
		}(i)
	}

	for {
		select {
		case <-quit:
			//закрываем канал
			close(quitGorutine)
			wg.Wait()
			log.Println("Process shutdown")
			return
		case <-time.After(time.Second):
			// шлем в канал случайные числа
			ch <- rand.Intn(100)
		}
	}
}
