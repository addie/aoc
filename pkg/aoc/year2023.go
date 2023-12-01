package aoc

import (
	"strconv"
	"strings"
	"unicode"
)

func (s Solution[T]) Year2023Day1(_ string) (int, int) {
	data := ReadFile("data/year2023day1.txt")
	res1 := year2023Day1Part1(data)
	res2 := year2023Day1Part2(data)
	return res1, res2
}

func year2023Day1Part1(data []string) int {
	res := 0
	for _, line := range data {
		left := ""
		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				left = string(line[i])
				break
			}
		}
		right := ""
		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				right = string(line[i])
				break
			}
		}
		intVal, _ := strconv.Atoi(left + right)
		res += intVal
	}
	return res
}

func year2023Day1Part2(data []string) int {
	var numMap = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	res := 0
	for _, line := range data {
		digits := findDigits(numMap, line)
		curr := digits[0]*10 + digits[len(digits)-1]
		res += curr
	}
	return res
}

func findDigits(numMap map[string]int, line string) []int {
	var digits []int
	for i, char := range line {
		if unicode.IsDigit(char) {
			val, _ := strconv.Atoi(string(char))
			digits = append(digits, val)
			continue
		}
		for k := range numMap {
			if strings.HasPrefix(line[i:], k) {
				digits = append(digits, numMap[k])
			}
		}
	}
	return digits
}
