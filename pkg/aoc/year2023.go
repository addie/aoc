package aoc

import (
	"fmt"
	"log"
	"strconv"
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
	var numMap = map[string]string{
		"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
		"six": "6", "seven": "7", "eight": "8", "nine": "9",
	}
	trie1 := NewTrie()
	for num := range numMap {
		trie1.Insert(num)
	}
	trie2 := NewTrie()
	for num := range numMap {
		trie2.InsertReversed(num)
	}
	res := 0
	for _, line := range data {
		left := ""
		var soFar string
		for i := 0; i < len(line); i++ {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				left = string(line[i])
				break
			}
			soFar += string(line[i])
			isValid := trie1.StartsWith(soFar)
			isComplete := trie1.Search(soFar)
			if !isValid {
				soFar = string(line[i])
				continue
			}
			if isComplete {
				left = numMap[soFar]
				break
			}
		}
		right := ""
		soFar = ""
		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				right = string(line[i])
				break
			}
			soFar += string(line[i])
			isValid := trie2.StartsWith(soFar)
			isComplete := trie2.Search(soFar)
			if !isValid {
				soFar = string(line[i])
				continue
			}
			if isComplete {
				soFar = Reverse(soFar)
				right = numMap[soFar]
				break
			}
		}
		intVal, err := strconv.Atoi(left + right)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(line)
		fmt.Println(intVal)
		res += intVal
	}
	return res
}
