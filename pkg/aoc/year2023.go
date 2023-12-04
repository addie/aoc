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

func isDigit(s string) bool { return strings.ContainsAny(s, "0123456789") }

func year2023Day3Part1(data []string) int {
	grid := make([][]string, len(data))
	for i := range data {
		grid[i] = strings.Split(data[i], "")
	}
	var partNumbers []int
	isPrevDigit := false
	isPartNo := false
	currDigit := ""
	checkPartNumber := func(grid [][]string, r, c int) bool {
		directions := []coord{{r: 0, c: 1}, {r: 1, c: 0}, {r: 1, c: 1}, {r: -1, c: 0}, {r: 0, c: -1}, {r: -1, c: -1}, {r: -1, c: 1}, {r: 1, c: -1}}
		for _, dir := range directions {
			nc := coord{r: r + dir.r, c: c + dir.c}
			if nc.r < 0 || nc.r >= len(grid) || nc.c < 0 || nc.c >= len(grid[0]) {
				continue
			}
			if grid[nc.r][nc.c] != "." && !isDigit(grid[nc.r][nc.c]) {
				return true
			}
		}
		return false
	}
	nextIsDigit := func(grid [][]string, r, c int) bool {
		if c+1 < len(grid[0]) && isDigit(grid[r][c+1]) {
			return true
		}
		return false
	}
	for r := range grid {
		for c := range grid[r] {
			cell := grid[r][c]
			if isDigit(cell) {
				currDigit += cell
			} else {
				currDigit = ""
				isPrevDigit = false
				isPartNo = false
				continue
			}
			if !nextIsDigit(grid, r, c) {
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				if isPartNo {
					partNumbers = append(partNumbers, Must(strconv.Atoi(currDigit)))
					currDigit = ""
					continue
				}
			}
			if !isPrevDigit {
				isPrevDigit = true
				currDigit = cell
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				continue
			} else {
				if checkPartNumber(grid, r, c) {
					isPartNo = true
				}
				continue
			}
		}
		currDigit = ""
		isPrevDigit = false
		isPartNo = false
	}
	sumOfPartNos := 0
	for _, partNo := range partNumbers {
		sumOfPartNos += partNo
	}
	return sumOfPartNos
}

func year2023Day3Part2(data []string) int {
	grid := make([][]string, len(data))
	for i := range data {
		grid[i] = strings.Split(data[i], "")
	}
	R, C := len(grid), len(grid[0])
	nums := make(map[coord][]int)
	for r := range grid {
		// set containing positions of '*' (gears)
		// collect all the gears surrounding each number group
		gears := make(map[coord]struct{})
		n := 0
		for c := range grid[0] {
			if c < C && isDigit(grid[r][c]) {
				n = n*10 + Must(strconv.Atoi(grid[r][c]))
				for _, rd := range []int{-1, 0, 1} {
					for _, cd := range []int{-1, 0, 1} {
						nr, nc := r+rd, c+cd
						if nr >= 0 && nr < R && nc >= 0 && nc < C {
							newRC := grid[nr][nc]
							if newRC == "*" {
								gears[coord{r: nr, c: nc}] = struct{}{}
							}
						}
					}
				}
			} else if n > 0 { // every time you pass a number, add the gears to the map
				for gear := range gears {
					nums[gear] = append(nums[gear], n)
				}
				n = 0
				gears = make(map[coord]struct{})
			}
		}
	}
	res := 0
	for _, v := range nums {
		if len(v) == 2 {
			res += v[0] * v[1]
		}
	}
	return res
}

func (s Solution[T]) Year2023Day4(_ string) (int, int) {
	data := ReadFile("data/year2023day4.txt")
	res1 := year2023Day3Part1(data)
	res2 := year2023Day3Part2(data)
	return res1, res2
}
