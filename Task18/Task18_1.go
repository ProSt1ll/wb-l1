package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//создаем структуру счетчика

type Counter struct {
	count int
	mutex sync.Mutex
}

//потокобезопасно инкрементируем
func (c *Counter) increment() {
	defer c.mutex.Unlock()
	c.mutex.Lock()
	c.count++

}

func main() {
	//задаем количество воркеров
	n := 10
	//канал для завершения
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//вейтгруппа для синхронизации
	wg := sync.WaitGroup{}
	//экземпляр структуры
	count := Counter{}
	//контекст для завершения
	ctx, cancel := context.WithCancel(context.Background())
	//создаем горутинки в цикле
	for i := 0; i < n; i++ {
		wg.Add(1)
		//слушаем завершение и инкрементируем счетчик
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					count.increment()
					time.Sleep(time.Second)
				}
			}
		}()
	}
	//флаг для выхода из вечного цикла
	var flag = true
	//выходим когда получаем сигнал
	for flag {
		select {
		case <-quit:
			flag = false
		}
	}
	//завершаем горутинки
	cancel()
	//ждем всех
	wg.Wait()
	//вывод результа
	fmt.Println(count.count)
}
