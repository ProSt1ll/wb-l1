package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//сделаем обертку в мапу, используем RWMutex тк подразумеваем, что читать будут намного чаще, чем писать
// это дает приемущетсво и быстродействие при чтении

type SafeMap struct {
	M     map[int]int
	mutex sync.RWMutex
}

func New() *SafeMap {
	return &SafeMap{
		M:     make(map[int]int),
		mutex: sync.RWMutex{},
	}
}

func (sm *SafeMap) Load(key int) (int, bool) {
	//Локаем на запись и читаем
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()
	val, ok := sm.M[key]
	return val, ok
}

func (sm *SafeMap) Store(key int, value int) {
	//локаем полностью и читаем
	sm.mutex.Lock()
	defer sm.mutex.Unlock()
	sm.M[key] = value
}

func main() {
	//вспомогательное число для установки максиального индекса мапы и кол-ва воркеров
	n := 10
	//сонтекст для остановки горутин
	ctx, cancel := context.WithCancel(context.Background())
	//экземпляр структуры
	sm := New()
	//создаем вейтгруппу для синхронизации(чтобы процесс раньше горутинки не завершился)
	wg := &sync.WaitGroup{}
	//в цикле добавляем горутинки
	for i := 0; i < n; i++ {
		wg.Add(1)
		//функция пишет в мапу и ждет завершения
		go func() {
			//устанавливаем defer функцию выполненной операции(декремент счетчика) группы ожидания
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
	for key, val := range sm.M {
		fmt.Println(key, ":", val)
	}

}
