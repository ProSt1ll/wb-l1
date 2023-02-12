package main

import (
	"fmt"
	"sync"
)

func main() {
	var input = []int{2, 4, 6, 8, 10}

	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := sync.WaitGroup{}

	//канал для передачи данных
	ch := make(chan int)

	//канал для завершения горутин
	quit := make(chan struct{})

	//переменная выхода
	var output int

	//анонимная функция для обработки каналов
	go func() {
		for {
			select {
			//если пришло по каналу данных, обрабатываем
			case a := <-ch:
				output += a * a
			//если по каналу завершения, закрываем горутинку
			case <-quit:
				return
			}
		}
	}()

	wg.Add(1)
	//анонимная функция для записи данных в канал
	go func() {
		//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
		defer wg.Done()
		//в цикле пишем в канал
		for _, val := range input {
			ch <- val
		}
		//выходим из горутинки обработки
		quit <- struct{}{}

	}()

	//ждем все горутинки
	wg.Wait()

	//результат
	fmt.Print(output)
}