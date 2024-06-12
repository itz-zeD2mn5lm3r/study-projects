package queue

import "fmt"

type Queue struct {
	s           []any // слайс в котором хранятся значения
	low, high   int   // индексы верхней и нижней границы очереди
	size, count int   // размер очереди и количество элементов
}

func NewQueue(size int) *Queue {
	return &Queue{
		s:    make([]any, size),
		size: size,
		low:  0,
		high: -1,
	}
}

// Push - добавление в очередь значения
func (q *Queue) Push(v any) {
	if q.count == q.size {
		fmt.Println("Очередь переполнена")
		return
	}
	q.high = (q.high + 1) % q.size
	q.s[q.high] = v
	q.count++
}

// Pop - получение значения из очереди и его удаление
func (q *Queue) Pop() any {
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
