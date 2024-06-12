package main

import (
	"fmt"
	"queue/queuepack"
)

// Переделать прошлое домашнее задание:
// вынести структуры данных в отдельные пакеты
// переделать функции которые обрабатывают структуры данных в методы

func main() {
	q := queue.NewQueue(5)
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop()) // 1
	fmt.Println(q.Pop()) // 2
	fmt.Println(q.Pop()) // 3
	fmt.Println(q.Pop()) // Очередь пуста, nil
}
