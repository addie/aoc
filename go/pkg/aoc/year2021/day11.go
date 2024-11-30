package aoc

import (
	"bufio"
	"os"
	"strconv"
)

const day11Filename = "data/day11"

func Day11Part1() int {
	file, _ := os.Open(day11Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		strRow := scanner.Text()
		var intRow []int
		for _, c := range strRow {
			ch, _ := strconv.Atoi(string(c))
			intRow = append(intRow, ch)
		}
		grid = append(grid, intRow)
	}
	const steps = 100
	flashes := 0
	for i := 0; i < steps; i++ {
		flashes += modelStep(grid)
	}
	return flashes
}

func modelStep(grid [][]int) int {
	flashes := 0

	for r := range grid {
		for c := range grid[0] {
			grid[r][c] += 1
		}
	}

	flashed := initFlashed(grid)
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] > 9 && !flashed[r][c] {
				flashed[r][c] = true
				bumpNeighbors(grid, r, c, flashed)
			}
		}
	}

	for r := range flashed {
		for c := range flashed[0] {
			if flashed[r][c] {
				flashes++
				grid[r][c] = 0
			}
		}
	}
	return flashes
}

func initFlashed(grid [][]int) [][]bool {
	flashed := make([][]bool, len(grid))
	for r := range grid {
		flashed[r] = make([]bool, len(grid[0]))
	}
	return flashed
}

var dirs = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {-1, -1}, {-1, 1}, {1, -1}}

func bumpNeighbors(grid [][]int, r int, c int, flashed [][]bool) {
	for _, dir := range dirs {
		nR, nC := r+dir[0], c+dir[1]
		if nR < 0 || nR >= len(grid) || nC < 0 || nC >= len(grid[0]) {
			continue
		}
		grid[nR][nC]++
		if grid[nR][nC] > 9 && !flashed[nR][nC] {
			flashed[nR][nC] = true
			bumpNeighbors(grid, nR, nC, flashed)
		}
	}
}

func Day11Part2() int {
	file, _ := os.Open(day11Filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]int
	for scanner.Scan() {
		strRow := scanner.Text()
		var intRow []int
		for _, c := range strRow {
			ch, _ := strconv.Atoi(string(c))
			intRow = append(intRow, ch)
		}
		grid = append(grid, intRow)
	}
	step := 0
	for {
		step++
		modelStep(grid)
		if allFlashed(grid) {
			return step
		}
	}
	return 0
}

func allFlashed(grid [][]int) bool {
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] != 0 {
				return false
			}
		}
	}
	return true
}
