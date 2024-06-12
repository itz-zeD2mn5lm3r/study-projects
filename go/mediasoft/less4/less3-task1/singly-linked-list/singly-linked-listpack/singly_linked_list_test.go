package linkedlist

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	l := NewSinglyLinkedList()
	if l.size != 0 {
		t.Errorf("Ожидалось, что размер будет 0, получено %d", l.size)
	}
	if l.first != nil {
		t.Error("Ожидалось, что first будет nil")
	}
	if l.last != nil {
		t.Error("Ожидалось, что last будет nil")
	}
}

func TestAdd(t *testing.T) {
	l := NewSinglyLinkedList()
	l.Add(1)
	if l.size != 1 {
		t.Errorf("Ожидалось, что размер будет 1, получено %d", l.size)
	}
	if l.first == nil || l.first.Value != 1 {
		t.Error("Ожидалось, что first будет 1")
	}
	if l.last == nil || l.last.Value != 1 {
		t.Error("Ожидалось, что last будет 1")
	}

	l.Add(2)
	if l.size != 2 {
		t.Errorf("Ожидалось, что размер будет 2, получено %d", l.size)
	}
	if l.last.Value != 2 {
		t.Errorf("Ожидалось, что last будет 2, получено %v", l.last.Value)
	}
}

func TestGet(t *testing.T) {
	l := NewSinglyLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	v := l.Get(1)
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
}

func TestRemove(t *testing.T) {
	l := NewSinglyLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Remove(1)
	if l.size != 2 {
		t.Errorf("Ожидалось, что размер будет 2, получено %d", l.size)
	}
	if l.Get(1) != 3 {
		t.Errorf("Ожидалось, что значение на индексе 1 будет 3, получено %v", l.Get(1))
	}

	l.Remove(0)
	if l.size != 1 {
		t.Errorf("Ожидалось, что размер будет 1, получено %d", l.size)
	}
	if l.Get(0) != 3 {
		t.Errorf("Ожидалось, что значение на индексе 0 будет 3, получено %v", l.Get(0))
	}

	l.Remove(0)
	if l.size != 0 {
		t.Errorf("Ожидалось, что размер будет 0, получено %d", l.size)
	}
	if l.first != nil {
		t.Error("Ожидалось, что first будет nil")
	}
	if l.last != nil {
		t.Error("Ожидалось, что last будет nil")
	}
}

func TestValues(t *testing.T) {
	l := NewSinglyLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	vals := l.Values()
	expected := []any{1, 2, 3}
	for i, v := range vals {
		if v != expected[i] {
			t.Errorf("Ожидалось, что значение на индексе %d будет %v, получено %v", i, expected[i], v)
		}
	}
}
