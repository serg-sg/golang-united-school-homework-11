package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	// n	- количество пользователей
	// pool	- количество горутин
	// res	- массив полученных пользователей

	// Инициализируем массив результата ID пользователей
	res = make([]user, n)

	// Определим переменную группу горутин
	var wg sync.WaitGroup

	// Создадим и инициализируем канал для сигнала завершения работы
	ch := make(chan struct{}, pool)

	// Цикл заполнения res
	for i := int64(0); i < n; i++ {
		// Добавим группу из 1 горутины
		wg.Add(1)
		ch <- struct{}{}

		// Запускаем анонимную функцию с передачей внутрь неё копии i
		go func(i int64) {
			defer wg.Done() // Отложенная функция отправки сигнала завершения работы
			res[i] = getOne(i)
			<-ch
		}(i)
	}
	// Ожидание завершения работы всех горутин
	wg.Wait()

	return res
}
