package main

import (
	"fmt"
)

func main() {
	// Nil канал
	var nilChan chan int

	// Закрытый канал
	closedChan := make(chan int)
	close(closedChan)

	// Не nil и не закрытый канал
	notNilNotClosedChan := make(chan int)

	// Канал для синхронизации завершения горутин
	done := make(chan struct{})

	// Случай 1: close(nilChan) - паника
	// Раскомментирование следующей строки вызовет панику
	// close(nilChan)

	// Случай 2: close(closedChan) - паника
	// Раскомментирование следующей строки вызовет панику
	// close(closedChan)

	// Случай 3: close(notNilNotClosedChan) - закрытие канала
	go func() {
		// Задержка для имитации работы
		notNilNotClosedChan <- 1
		close(notNilNotClosedChan)
		close(done) // Уведомляем о завершении горутины
	}()

	// Случай 4: nilChan <- val - заблокируется навсегда
	// Раскомментирование следующей строки приведет к блокировке горутины навсегда
	// go func() { nilChan <- 1 }()

	// Случай 5: closedChan <- val - паника
	// Раскомментирование следующей строки вызовет панику
	// closedChan <- 1

	// Случай 6: notNilNotClosedChan <- val - посылаем или блокируемся
	// Убрано для предотвращения ошибки

	// Случай 7: val := <-nilChan - заблокируется навсегда
	// Раскомментирование следующей строки приведет к блокировке горутины навсегда
	// go func() { <-nilChan }()

	// Случай 8: val := <-closedChan - не блокирует
	val, ok := <-closedChan
	fmt.Printf("Случай 8: Чтение из закрытого канала - значение: %v, ok: %v\n", val, ok)

	// Случай 9: val := <-notNilNotClosedChan - читаем или блокируемся
	val = <-notNilNotClosedChan
	fmt.Printf("Случай 9: Чтение из не nil и не закрытого канала - значение: %v\n", val)

	// Ожидаем завершения всех горутин
	<-done

	// Используем nilChan, чтобы избежать ошибки о неиспользуемой переменной
	_ = nilChan
}
