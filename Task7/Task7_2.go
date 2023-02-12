package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//https://habr.com/ru/post/338718/
//статья на хабре о том, что синк мап бесполезнее если у нас не высонагруженный сервис, где решают 100нс и при этом
//имеем 32+ ядерный процессор

func main() {
	//вспомогательное число для установки максиального индекса мапы и кол-ва воркеров
	n := 10
	//сонтекст для остановки горутин
	ctx, cancel := context.WithCancel(context.Background())

	// sync.Map: - я для вас шутка какая то?
	sm := sync.Map{}

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}
	//в цикле добавляем горутинки
	for i := 0; i < n; i++ {
		wg.Add(1)
		//функция пишет в мапу и ждет завершения
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					sm.Store(rand.Intn(n), rand.Intn(100))

				}
			}
		}()

	}
	//немного ждем
	time.Sleep(time.Second * 3)

	//завершаем
	cancel()

	//ждем завершения
	wg.Wait()
	//вывод
	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
