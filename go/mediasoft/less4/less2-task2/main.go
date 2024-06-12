package main

import (
	"fmt"
	"strings"
)

// Условие задачи:
// Сделать конвертер из римских цифр в арабские

// romanToInt - конвертирует римское число в арабское
func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	total := roman[s[n-1]]

	for i := n - 2; i >= 0; i-- {
		if roman[s[i]] < roman[s[i+1]] {
			total -= roman[s[i]]
		} else {
			total += roman[s[i]]
		}
	}
	return total
}

func main() {
	var romanNumeral string
	fmt.Print("Введите римское число: ")
	if _, err := fmt.Scanln(&romanNumeral); err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	romanNumeral = strings.TrimSpace(romanNumeral)

	arabicNumber := romanToInt(romanNumeral)
	fmt.Printf("Арабское число для %s равно %d\n", romanNumeral, arabicNumber)
}
