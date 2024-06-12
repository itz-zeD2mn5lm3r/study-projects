package main

import (
	"testing"
)

// TestQueue - тестирование основных функций очереди
func TestQueue(t *testing.T) {
	q := newQueue(3)

	// Тестирование извлечения элемента из пустой очереди
	if v := pop(q); v != nil {
		t.Errorf("ожидалось nil, получено %v", v)
	}

	// Тестирование добавления элементов в очередь
	push(q, 1)
	push(q, 2)
	push(q, 3)

	// Тестирование переполнения очереди
	push(q, 4)
	if q.count != 3 {
		t.Errorf("ожидалось 3 элемента, получено %d", q.count)
	}

	// Тестирование извлечения элементов из очереди
	if v := pop(q); v != 1 {
		t.Errorf("ожидалось 1, получено %v", v)
	}
	if v := pop(q); v != 2 {
		t.Errorf("ожидалось 2, получено %v", v)
	}
	if v := pop(q); v != 3 {
		t.Errorf("ожидалось 3, получено %v", v)
	}

	// Тестирование извлечения элемента из пустой очереди
	if v := pop(q); v != nil {
		t.Errorf("ожидалось nil, получено %v", v)
	}
}

// TestQueueWrapAround - тестирование циклического заполнения очереди
func TestQueueWrapAround(t *testing.T) {
	q := newQueue(3)

	// Тестирование циклического заполнения
	push(q, 1)
	push(q, 2)
	push(q, 3)

	if v := pop(q); v != 1 {
		t.Errorf("ожидалось 1, получено %v", v)
	}
	push(q, 4)
	if v := pop(q); v != 2 {
		t.Errorf("ожидалось 2, получено %v", v)
	}
	if v := pop(q); v != 3 {
		t.Errorf("ожидалось 3, получено %v", v)
	}
	if v := pop(q); v != 4 {
		t.Errorf("ожидалось 4, получено %v", v)
	}
}

// TestQueueOverflow - тестирование переполнения очереди
func TestQueueOverflow(t *testing.T) {
	q := newQueue(3)

	// Заполнение очереди
	push(q, 1)
	push(q, 2)
	push(q, 3)

	// Попытка добавления элемента в переполненную очередь
	push(q, 4)
	if q.count != 3 {
		t.Errorf("ожидалось 3 элемента, получено %d", q.count)
	}

	// Убедиться, что очередь работает корректно после переполнения
	if v := pop(q); v != 1 {
		t.Errorf("ожидалось 1, получено %v", v)
	}
	if v := pop(q); v != 2 {
		t.Errorf("ожидалось 2, получено %v", v)
	}
	if v := pop(q); v != 3 {
		t.Errorf("ожидалось 3, получено %v", v)
	}
	if v := pop(q); v != nil {
		t.Errorf("ожидалось nil, получено %v", v)
	}
}
