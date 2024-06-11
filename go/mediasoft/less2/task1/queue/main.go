package main

import "fmt"

// Условие задачи:
// Реализуйте структуру данных 'очередь' с функциями инициализации, добавления элемента и извлечения элемента.

type queue struct {
	s           []any // слайс в котором хранятся значения
	low, high   int   // индексы верхней и нижней границы очереди
	size, count int   // размер очереди и количество элементов
}

func newQueue(size int) *queue {
	return &queue{
		s:    make([]any, size),
		size: size,
		low:  0,
		high: -1,
	}
}

// push - добавление в очередь значения
func push(q *queue, v any) {
	if q.count == q.size {
		fmt.Println("Очередь переполнена")
		return
	}
	q.high = (q.high + 1) % q.size
	q.s[q.high] = v
	q.count++
}

// pop - получение значения из очереди и его удаление
func pop(q *queue) any {
	if q.count == 0 {
		fmt.Println("Очередь пуста")
		return nil
	}
	v := q.s[q.low]
	q.s[q.low] = nil
	q.low = (q.low + 1) % q.size
	q.count--
	return v
}

func main() {
	q := newQueue(5)
	push(q, 1)
	push(q, 2)
	push(q, 3)
	fmt.Println(pop(q)) // 1
	fmt.Println(pop(q)) // 2
	fmt.Println(pop(q)) // 3
	fmt.Println(pop(q)) // Очередь пуста, nil
}
