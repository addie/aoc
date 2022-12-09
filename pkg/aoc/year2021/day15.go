package aoc

import (
	"aoc2021/aoc/new"
	"aoc2021/clipboard"
	"bufio"
	"os"
	"strconv"
	"strings"
)

const day15Filename = "data/day15"
const gridSize = 100
const memoSize = 500
const section = memoSize / 5

func Day15() int {
	file, _ := os.Open(day15Filename)
	defer file.Close()

	grid := new.IntGrid(gridSize)
	row := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		r := strings.TrimSpace(scanner.Text())
		for col := range r {
			val, _ := strconv.Atoi(string(r[col]))
			grid[row][col] = val
		}
		row++
	}

	memo := new.IntGrid(memoSize)
	for c := 1; c < memoSize; c++ {
		gridValue := calcValue(grid, 0, c)
		memo[0][c] = gridValue + memo[0][c-1]
	}
	for r := 1; r < memoSize; r++ {
		gridValue := calcValue(grid, r, 0)
		memo[r][0] = gridValue + memo[r-1][0]
	}
	for r := 1; r < memoSize; r++ {
		for c := 1; c < memoSize; c++ {
			gridValue := calcValue(grid, r, c)
			memo[r][c] = gridValue + minPath(memo[r][c-1], memo[r-1][c])
		}
	}
	var strs, allStrs []string
	for _, r := range grid {
		for _, c := range r {
			s := strconv.Itoa(c)
			strs = append(strs, s)
		}
		allStrs = append(allStrs, strings.Join(strs, ","))
	}
	res := strings.Join(allStrs, "\n")
	clipboard.WriteAll(res)
	return memo[len(memo)-1][len(memo[0])-1]
}

func calcValue(grid [][]int, r int, c int) int {
	val := grid[r%section][c%section]
	loop := c/section + r/section
	for loop > 0 {
		val++
		if val > 9 {
			val = 1
		}
		loop--
	}
	return val
}

func minPath(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}
