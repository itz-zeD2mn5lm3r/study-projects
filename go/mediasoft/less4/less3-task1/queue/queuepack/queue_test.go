package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(3)
	if q.size != 3 {
		t.Errorf("Ожидалось, что размер очереди будет 3, получено %d", q.size)
	}
	if q.low != 0 {
		t.Errorf("Ожидалось, что low будет 0, получено %d", q.low)
	}
	if q.high != -1 {
		t.Errorf("Ожидалось, что high будет -1, получено %d", q.high)
	}
	if q.count != 0 {
		t.Errorf("Ожидалось, что count будет 0, получено %d", q.count)
	}
}

func TestPush(t *testing.T) {
	q := NewQueue(2)
	q.Push(1)
	if q.s[q.high] != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", q.s[q.high])
	}
	if q.count != 1 {
		t.Errorf("Ожидалось, что count будет 1, получено %d", q.count)
	}

	q.Push(2)
	if q.s[q.high] != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", q.s[q.high])
	}
	if q.count != 2 {
		t.Errorf("Ожидалось, что count будет 2, получено %d", q.count)
	}

	q.Push(3)
	if q.count != 2 {
		t.Errorf("Ожидалось, что count будет 2, получено %d", q.count)
	}
}

func TestPop(t *testing.T) {
	q := NewQueue(2)
	q.Push(1)
	q.Push(2)

	v := q.Pop()
	if v != 1 {
		t.Errorf("Ожидалось, что значение будет 1, получено %v", v)
	}
	if q.count != 1 {
		t.Errorf("Ожидалось, что count будет 1, получено %d", q.count)
	}

	v = q.Pop()
	if v != 2 {
		t.Errorf("Ожидалось, что значение будет 2, получено %v", v)
	}
	if q.count != 0 {
		t.Errorf("Ожидалось, что count будет 0, получено %d", q.count)
	}

	v = q.Pop()
	if v != nil {
		t.Errorf("Ожидалось, что значение будет nil, получено %v", v)
	}
	if q.count != 0 {
		t.Errorf("Ожидалось, что count будет 0, получено %d", q.count)
	}
}
