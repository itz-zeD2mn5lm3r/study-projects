package converter

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman  string
		arabic int
		err    error
	}{
		{"I", 1, nil},
		{"II", 2, nil},
		{"III", 3, nil},
		{"IV", 4, nil},
		{"V", 5, nil},
		{"IX", 9, nil},
		{"X", 10, nil},
		{"XL", 40, nil},
		{"L", 50, nil},
		{"XC", 90, nil},
		{"C", 100, nil},
		{"CD", 400, nil},
		{"D", 500, nil},
		{"CM", 900, nil},
		{"M", 1000, nil},
		{"MCMXCIV", 1994, nil},
		{"MMXXIII", 2023, nil},
		{"", 0, fmt.Errorf("empty input")},
	}

	for _, test := range tests {
		result, err := RomanToInt(test.roman)
		if result != test.arabic || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf("Ожидалось %d и %v для %s, получено %d и %v", test.arabic, test.err, test.roman, result, err)
		}
	}
}
