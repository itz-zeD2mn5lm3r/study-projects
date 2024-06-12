package main

import "fmt"

// Условие задачи:
// Реализуйте структуру данных 'односвязный список' с функциями инициализации, добавления элемента,
// получения элемента по индексу, удаления элемента по индексу и получения всех элементов в виде слайса.

type singlyLinkedList struct {
	first *item
	last  *item
	size  int
}

type item struct {
	v    any
	next *item
}

func newSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{}
}

// add - добавление значения в связный список
func add(l *singlyLinkedList, v any) {
	newItem := &item{v: v}
	if l.first == nil {
		l.first = newItem
		l.last = newItem
	} else {
		l.last.next = newItem
		l.last = newItem
	}
	l.size++
}

// get - получение значения по индексу из связанного списка
func get(l *singlyLinkedList, idx int) any {
	if idx < 0 || idx >= l.size {
		fmt.Println("Index out of bounds")
		return nil
	}
	current := l.first
	for i := 0; i < idx; i++ {
		current = current.next
	}
	return current.v
}

// remove - удаление значения по индексу из списка
func remove(l *singlyLinkedList, idx int) {
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

// values - получение слайса значений из списка
func values(l *singlyLinkedList) []any {
	vals := make([]any, 0, l.size)
	current := l.first
	for current != nil {
		vals = append(vals, current.v)
		current = current.next
	}
	return vals
}

func main() {
	l := newSinglyLinkedList()
	add(l, 1)
	add(l, 2)
	add(l, 3)
	fmt.Println(get(l, 1)) // 2
	remove(l, 1)
	fmt.Println(values(l)) // [1 3]
}
