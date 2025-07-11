package calculator

import (
	"strings"
)

func IntToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := ""

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result += syms[i]
		}
	}
	return result
}
func RomanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	prev := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := symbols[s[i]]
		if curr < prev {
			total -= curr
		} else {
			total += curr
		}
		prev = curr
	}

	return total
}
func IsValidRoman(s string) bool {
	for _, r := range s {
		if !strings.ContainsRune("IVXLCDM", r) {
			return false
		}
	}
	return true
}
