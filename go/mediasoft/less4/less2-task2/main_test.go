package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman  string
		arabic int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"X", 10},
		{"XL", 40},
		{"L", 50},
		{"XC", 90},
		{"C", 100},
		{"CD", 400},
		{"D", 500},
		{"CM", 900},
		{"M", 1000},
		{"MCMXCIV", 1994},
		{"MMXXIII", 2023},
	}

	for _, test := range tests {
		result := romanToInt(test.roman)
		if result != test.arabic {
			t.Errorf("Ожидалось %d для %s, получено %d", test.arabic, test.roman, result)
		}
	}
}
