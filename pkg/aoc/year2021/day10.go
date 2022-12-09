package aoc

import (
	"bufio"
	"os"
	"sort"
)

const day10Filename = "data/day10"

func Day10Part1() int {
	file, _ := os.Open(day10Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []rune{'[', '(', '{', '<'}
	var wrongs []rune
	for scanner.Scan() {
		row := scanner.Text()
		var stack []rune
		for _, c := range row {
			if inRune(left, c) {
				stack = append(stack, c)
			} else {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == '(' && (c == '}' || c == ']' || c == '>') ||
					top == '{' && (c == ')' || c == ']' || c == '>') ||
					top == '[' && (c == ')' || c == '}' || c == '>') ||
					top == '<' && (c == ')' || c == ']' || c == '}') {
					wrongs = append(wrongs, c)
				}
			}
		}
	}
	score := 0
	for _, wrong := range wrongs {
		if wrong == ')' {
			score += 3
		}
		if wrong == ']' {
			score += 57
		}
		if wrong == '}' {
			score += 1197
		}
		if wrong == '>' {
			score += 25137
		}
	}
	return score
}

func inRune(col []rune, c rune) bool {
	for _, item := range col {
		if c == item {
			return true
		}
	}
	return false
}

func Day10Part2() int {
	file, _ := os.Open(day10Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	left := []rune{'[', '(', '{', '<'}
	var scores []int
	var incompleteRows []string
	for scanner.Scan() {
		row := scanner.Text()
		incomplete := isIncompleteRow(row, left)
		if incomplete {
			incompleteRows = append(incompleteRows, row)
		}
	}
	for _, row := range incompleteRows {
		stack := completeRow(row)
		score := 0
		for i := len(stack) - 1; i >= 0; i-- {
			score *= 5
			if stack[i] == "(" {
				score += 1
			}
			if stack[i] == "[" {
				score += 2
			}
			if stack[i] == "{" {
				score += 3
			}
			if stack[i] == "<" {
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	mid := len(scores) / 2
	return scores[mid]
}

func completeRow(row string) []string {
	var s []string
	for _, c := range row {
		switch c {
		case '{', '[', '<', '(':
			s = append(s, string(c))
		case '}':
			s = s[:len(s)-1]
		case ']':
			s = s[:len(s)-1]
		case '>':
			s = s[:len(s)-1]
		case ')':
			s = s[:len(s)-1]
		}
	}
	return s
}

func isIncompleteRow(row string, left []rune) bool {
	var stack []rune
	for _, c := range row {
		if inRune(left, c) {
			stack = append(stack, c)
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top == '(' && (c == '}' || c == ']' || c == '>') ||
				top == '{' && (c == ')' || c == ']' || c == '>') ||
				top == '[' && (c == ')' || c == '}' || c == '>') ||
				top == '<' && (c == ')' || c == ']' || c == '}') {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return true
	}
	return false
}
