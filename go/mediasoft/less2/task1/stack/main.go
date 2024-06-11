package main

import "fmt"

// Условие задачи:
// Реализуйте структуру данных 'стек' с функциями инициализации, добавления, извлечения и просмотра элементов

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	if s.head+1 >= len(s.s) {
		fmt.Println("Переполнение стека")
		return
	}
	s.head++
	s.s[s.head] = v
}

// pop - получение значения из стека и его удаление из вершины
func pop(s *stack) any {
	if s.head == -1 {
		fmt.Println("Недостаток элементов в стеке")
		return nil
	}
	v := s.s[s.head]
	s.s[s.head] = nil
	s.head--
	return v
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head == -1 {
		fmt.Println("Переполнение стека")
		return nil
	}
	return s.s[s.head]
}

func main() {
	s := newStack(5)
	push(s, 1)
	push(s, 2)
	push(s, 3)
	fmt.Println(peek(s)) // 3
	fmt.Println(pop(s))  // 3
	fmt.Println(pop(s))  // 2
	fmt.Println(pop(s))  // 1
	fmt.Println(pop(s))  // Недостаток элементов в стеке, nil
}
