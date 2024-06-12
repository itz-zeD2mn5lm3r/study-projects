package stack

import "fmt"

type Stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func NewStack(size int) *Stack {
	return &Stack{
		s:    make([]any, size),
		head: -1,
	}
}

// Push - добавление в стек значения
func (s *Stack) Push(v any) {
	if s.head+1 >= len(s.s) {
		fmt.Println("Переполнение стека")
		return
	}
	s.head++
	s.s[s.head] = v
}

// Pop - получение значения из стека и его удаление из вершины
func (s *Stack) Pop() any {
	if s.head == -1 {
		fmt.Println("Недостаток элементов в стеке")
		return nil
	}
	v := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return v
}

// Peek - просмотр значения на вершине стека
func (s *Stack) Peek() any {
	if s.head == -1 {
		fmt.Println("Недостаток элементов в стеке")
		return nil
	}
	return s.s[s.head]
}
