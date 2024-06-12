package main

import (
	"fmt"
	"stack/stackpack"
)

// Переделать прошлое домашнее задание:
// вынести структуры данных в отдельные пакеты
// переделать функции которые обрабатывают структуры данных в методы

func main() {
	s := stack.NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek()) // 3
	fmt.Println(s.Pop())  // 3
	fmt.Println(s.Pop())  // 2
	fmt.Println(s.Pop())  // 1
	fmt.Println(s.Pop())  // Недостаток элементов в стеке, nil
}
