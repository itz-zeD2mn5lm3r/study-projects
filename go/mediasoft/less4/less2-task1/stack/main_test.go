package main

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := newStack(3)
	if s.head != -1 {
		t.Errorf("Ожидалось, что head будет -1, получено %d", s.head)
	}
	if len(s.s) != 3 {
		t.Errorf("Ожидалось, что размер стека будет 3, получено %d", len(s.s))
	}
}

func TestPush(t *testing.T) {
	s := newStack(2)
	push(s, 1)
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}
	if s.s[s.head] != 1 {
		t.Errorf("Ожидалось, что значение на вершине будет 1, получено %v", s.s[s.head])
	}

	push(s, 2)
	if s.head != 1 {
		t.Errorf("Ожидалось, что head будет 1, получено %d", s.head)
	}
	if s.s[s.head] != 2 {
		t.Errorf("Ожидалось, что значение на вершине будет 2, получено %v", s.s[s.head])
	}

	// Проверка переполнения стека
	push(s, 3)
	if s.head != 1 {
		t.Errorf("Ожидалось, что head останется 1, получено %d", s.head)
	}
}

func TestPop(t *testing.T) {
	s := newStack(2)
	push(s, 1)
	push(s, 2)

	v := pop(s)
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}

	v = pop(s)
	if v != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", v)
	}
	if s.head != -1 {
		t.Errorf("Ожидалось, что head будет -1, получено %d", s.head)
	}

	// Проверка недостатка элементов в стеке
	v = pop(s)
	if v != nil {
		t.Errorf("Ожидалось, что значение будет nil, получено %v", v)
	}
}

func TestPeek(t *testing.T) {
	s := newStack(2)
	push(s, 1)
	push(s, 2)

	v := peek(s)
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
	if s.head != 1 {
		t.Errorf("Ожидалось, что head будет 1, получено %d", s.head)
	}

	pop(s)
	v = peek(s)
	if v != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", v)
	}
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}

	pop(s)
	// Проверка недостатка элементов в стеке
	v = peek(s)
	if v != nil {
		t.Errorf("Ожидалось, что значение будет nil, получено %v", v)
	}
}
