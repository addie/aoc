package aoc

import (
	"bufio"
	"os"
	"strings"
)

const day0Filename = "data/day0"

func Day0() int {
	file, _ := os.Open(day16Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		_ = strings.TrimSpace(scanner.Text())
	}
	return 0
}
