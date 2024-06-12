package linkedlist

import "fmt"

type SinglyLinkedList struct {
	first *Item
	last  *Item
	size  int
}

type Item struct {
	Value any
	next  *Item
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Add - добавление значения в связный список
func (l *SinglyLinkedList) Add(v any) {
	newItem := &Item{Value: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// Get - получение значения по индексу из связанного списка
func (l *SinglyLinkedList) Get(idx int) any {
	if idx < 0 || idx >= l.size {
		fmt.Println("Index out of bounds")
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.Value
}

// Remove - удаление значения по индексу из списка
func (l *SinglyLinkedList) Remove(idx int) {
	if idx < 0 || idx >= l.size {
		fmt.Println("Index out of bounds")
		return
	}
	if idx == 0 {
		l.first = l.first.next
		if l.first == nil {
			l.last = nil
		}
	} else {
		previous := l.first
		for i := 0; i < idx-1; i++ {
			previous = previous.next
		}
		previous.next = previous.next.next
		if idx == l.size-1 {
			l.last = previous
		}
	}
	l.size--
}

// Values - получение слайса значений из списка
func (l *SinglyLinkedList) Values() []any {
	vals := make([]any, 0, l.size)
	current := l.first
	for current != nil {
		vals = append(vals, current.Value)
		current = current.next
	}
	return vals
}
