package aoc

import (
	"strconv"
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
	numMap, trie1, trie2 := createDataStructures()
	res := 0
	for _, line := range data {
		firstDigit := findDigit(numMap, trie1, line)
		lastDigit := findDigit(numMap, trie2, line, true)
		curr := firstDigit*10 + lastDigit
		res += curr
	}
	return res
}

func createDataStructures() (map[string]int, *Trie, *Trie) {
	var numMap = map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	trie := NewTrie()
	for num := range numMap {
		trie.Insert(num)
		trie.Insert(Reversed(num))
	}
	trie2 := NewTrie()
	for num := range numMap {
		trie2.InsertReversed(num)
	}
	return numMap, trie, trie2
}

func findDigit(numMap map[string]int, trie *Trie, line string, reversed ...bool) int {
	if reversed != nil && reversed[0] {
		line = Reversed(line)
	}
	digit := 0
	var soFar string
	for _, char := range line {
		if unicode.IsDigit(char) {
			val, _ := strconv.Atoi(string(char))
			return val
		}
		soFar += string(char)
		isValid := trie.StartsWith(soFar)
		isComplete := trie.Search(soFar)
		if isComplete {
			if reversed != nil && reversed[0] {
				soFar = Reversed(soFar)
			}
			return numMap[soFar]
		}
		if !isValid {
			soFar = string(char)
		}
	}
	return digit
}
