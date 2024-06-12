package main

import (
	"fmt"
	"singly-linked-list/singly-linked-listpack"
)

// Переделать прошлое домашнее задание:
// вынести структуры данных в отдельные пакеты
// переделать функции которые обрабатывают структуры данных в методы

func main() {
	l := linkedlist.NewSinglyLinkedList()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	fmt.Println(l.Get(1)) // 2
	l.Remove(1)
	fmt.Println(l.Values()) // [1 3]
}
