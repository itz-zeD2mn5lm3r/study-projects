package converter

import (
	"fmt"
	"strings"
)

// RomanToInt - функция конвертации римских цифр в арабские
func RomanToInt(s string) (int, error) {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0, fmt.Errorf("empty input")
	}

	n := len(s)
	total := romanMap[s[n-1]]

	for i := n - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			total -= romanMap[s[i]]
		} else {
			total += romanMap[s[i]]
		}
	}
	return total, nil
}
