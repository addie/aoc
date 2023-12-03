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

func (s Solution[T]) Year2023Day2(_ string) (int, int) {
	data := ReadFile("data/year2023day2.txt")
	res1 := year2023Day2Part1(data)
	res2 := year2023Day2Part2(data)
	return res1, res2
}

func year2023Day2Part1(data []string) int {
	type Color int

	const red Color = 12
	const green Color = 13
	const blue Color = 14

	var colorMap = map[string]Color{
		"red":   red,
		"green": green,
		"blue":  blue,
	}

	validGames := make([]bool, 0, len(data))
	for _, game := range data {
		gameValid := true
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, rd := range rounds {
			rd = strings.TrimSpace(rd)
			rolls := strings.Split(rd, ", ")
			for _, roll := range rolls {
				r := strings.Split(roll, " ")
				n, color := r[0], r[1]
				num, _ := strconv.Atoi(n)
				if num > int(colorMap[color]) {
					gameValid = false
					break
				}
			}
		}
		validGames = append(validGames, gameValid)
	}
	res := 0
	for i, valid := range validGames {
		if valid {
			res += i + 1
		}
	}
	return res
}

func year2023Day2Part2(data []string) int {
	powers := make([]int, 0, len(data))
	for _, game := range data {
		minBlue, minGreen, minRed := 0, 0, 0
		rounds := strings.Split(strings.Split(game, ":")[1], ";")
		for _, rd := range rounds {
			rd = strings.TrimSpace(rd)
			rolls := strings.Split(rd, ", ")
			for _, roll := range rolls {
				r := strings.Split(roll, " ")
				n, color := r[0], r[1]
				num, _ := strconv.Atoi(n)
				switch color {
				case "red":
					minRed = max(minRed, num)
				case "green":
					minGreen = max(minGreen, num)
				case "blue":
					minBlue = max(minBlue, num)
				}
			}
		}
		powers = append(powers, minRed*minBlue*minGreen)
	}
	res := 0
	for _, power := range powers {
		res += power
	}
	return res
}

func (s Solution[T]) Year2023Day3(_ string) (int, int) {
	data := ReadFile("data/year2023day3.txt")
	res1 := year2023Day3Part1(data)
	res2 := year2023Day3Part2(data)
	return res1, res2
}

func year2023Day3Part1(data []string) int {
	return 0
}

func year2023Day3Part2(data []string) int {
	return 0
}
