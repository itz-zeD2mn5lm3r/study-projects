package main

import (
	"testing"
)

func TestNewSinglyLinkedList(t *testing.T) {
	l := newSinglyLinkedList()
	if l.first != nil {
		t.Error("Ожидалось, что first будет nil")
	}
	if l.last != nil {
		t.Error("Ожидалось, что last будет nil")
	}
	if l.size != 0 {
		t.Error("Ожидалось, что size будет 0")
	}
}

func TestAdd(t *testing.T) {
	l := newSinglyLinkedList()
	add(l, 1)
	if l.first == nil {
		t.Error("Ожидалось, что first не будет nil")
	}
	if l.last == nil {
		t.Error("Ожидалось, что last не будет nil")
	}
	if l.size != 1 {
		t.Errorf("Ожидалось, что size будет 1, получено %d", l.size)
	}
	if l.first.v != 1 {
		t.Errorf("Ожидалось, что значение first будет 1, получено %v", l.first.v)
	}
}

func TestGet(t *testing.T) {
	l := newSinglyLinkedList()
	add(l, 1)
	add(l, 2)
	add(l, 3)
	val := get(l, 1)
	if val != 2 {
		t.Errorf("Ожидалось, что значение на индексе 1 будет 2, получено %v", val)
	}
}

func TestRemove(t *testing.T) {
	l := newSinglyLinkedList()
	add(l, 1)
	add(l, 2)
	add(l, 3)
	remove(l, 1)
	if l.size != 2 {
		t.Errorf("Ожидалось, что size будет 2, получено %d", l.size)
	}
	if l.first.next.v != 3 {
		t.Errorf("Ожидалось, что значение на индексе 1 будет 3, получено %v", l.first.next.v)
	}
	remove(l, 0)
	if l.size != 1 {
		t.Errorf("Ожидалось, что size будет 1, получено %d", l.size)
	}
	if l.first.v != 3 {
		t.Errorf("Ожидалось, что значение first будет 3, получено %v", l.first.v)
	}
	remove(l, 0)
	if l.size != 0 {
		t.Errorf("Ожидалось, что size будет 0, получено %d", l.size)
	}
	if l.first != nil {
		t.Error("Ожидалось, что first будет nil")
	}
	if l.last != nil {
		t.Error("Ожидалось, что last будет nil")
	}
}

func TestValues(t *testing.T) {
	l := newSinglyLinkedList()
	add(l, 1)
	add(l, 2)
	add(l, 3)
	vals := values(l)
	expected := []any{1, 2, 3}
	for i, v := range vals {
		if v != expected[i] {
			t.Errorf("Ожидалось, что значение на индексе %d будет %v, получено %v", i, expected[i], v)
		}
	}
}
