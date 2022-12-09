package aoc

import (
	"bufio"
	"os"
	"strings"
)

const day16Filename = "data/day16"

func Day16() int {
	file, _ := os.Open(day16Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = strings.TrimSpace(scanner.Text())
	}
	return 0
}
