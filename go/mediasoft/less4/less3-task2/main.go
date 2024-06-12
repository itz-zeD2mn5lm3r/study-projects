package main

import (
	"fmt"
	"less3-task2/converter"
)

// Переделать прошлое домашнее задание:
// вынести структуры данных в отдельные пакеты
// переделать функции которые обрабатывают структуры данных в методы

func main() {
	var romanNumeral string
	fmt.Print("Введите римское число: ")
	if _, err := fmt.Scanln(&romanNumeral); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	arabicNumber, err := converter.RomanToInt(romanNumeral)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	fmt.Printf("Арабское число для %s равно %d\n", romanNumeral, arabicNumber)
}
