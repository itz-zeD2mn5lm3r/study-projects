package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack(3)
	if s.head != -1 {
		t.Errorf("Ожидалось, что head будет -1, получено %d", s.head)
	}
	if len(s.s) != 3 {
		t.Errorf("Ожидалось, что размер стека будет 3, получено %d", len(s.s))
	}
}

func TestPush(t *testing.T) {
	s := NewStack(2)
	s.Push(1)
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}
	if s.s[s.head] != 1 {
		t.Errorf("Ожидалось, что значение на вершине будет 1, получено %v", s.s[s.head])
	}

	s.Push(2)
	if s.head != 1 {
		t.Errorf("Ожидалось, что head будет 1, получено %d", s.head)
	}
	if s.s[s.head] != 2 {
		t.Errorf("Ожидалось, что значение на вершине будет 2, получено %v", s.s[s.head])
	}

	// Проверка переполнения стека
	s.Push(3)
	if s.head != 1 {
		t.Errorf("Ожидалось, что head останется 1, получено %d", s.head)
	}
}

func TestPop(t *testing.T) {
	s := NewStack(2)
	s.Push(1)
	s.Push(2)

	v := s.Pop()
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}

	v = s.Pop()
	if v != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", v)
	}
	if s.head != -1 {
		t.Errorf("Ожидалось, что head будет -1, получено %d", s.head)
	}

	// Проверка недостатка элементов в стеке
	v = s.Pop()
	if v != nil {
		t.Errorf("Ожидалось, что значение будет nil, получено %v", v)
	}
}

func TestPeek(t *testing.T) {
	s := NewStack(2)
	s.Push(1)
	s.Push(2)

	v := s.Peek()
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
	if s.head != 1 {
		t.Errorf("Ожидалось, что head будет 1, получено %d", s.head)
	}

	s.Pop()
	v = s.Peek()
	if v != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", v)
	}
	if s.head != 0 {
		t.Errorf("Ожидалось, что head будет 0, получено %d", s.head)
	}

	s.Pop()
	// Проверка недостатка элементов в стеке
	v = s.Peek()
	if v != nil {
		t.Errorf("Ожидалось, что значение будет nil, получено %v", v)
	}
}
