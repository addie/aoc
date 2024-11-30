package aoc

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func Day9Part1() int {
	file, _ := os.Open("data/day9")
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		var r []string
		for _, c := range row {
			r = append(r, string(c))
		}
		grid = append(grid, r)
	}

	sumOfRiskLevels := 0
	for r := range grid {
		for c := range grid[0] {
			isLowPoint := true
			if r > 0 && grid[r-1][c] <= grid[r][c] {
				isLowPoint = false
			}
			if c > 0 && grid[r][c-1] <= grid[r][c] {
				isLowPoint = false
			}
			if r < len(grid)-1 && grid[r+1][c] <= grid[r][c] {
				isLowPoint = false
			}
			if c < len(grid[0])-1 && grid[r][c+1] <= grid[r][c] {
				isLowPoint = false
			}
			if isLowPoint {
				cellNum, _ := strconv.Atoi(grid[r][c])
				riskLevel := cellNum + 1
				sumOfRiskLevels += riskLevel
			}
		}
	}

	return sumOfRiskLevels
}

func Day9Part2() int {
	file, _ := os.Open("data/day9")
	defer file.Close()

	var grid [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		var r []string
		for _, c := range row {
			r = append(r, string(c))
		}
		grid = append(grid, r)
	}

	var basinSizes []int
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] != "9" {
				basinSizes = append(basinSizes, measureBasin(grid, r, c))
			}
		}
	}
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func measureBasin(grid [][]string, r int, c int) int {
	size := 0
	stack := []coord{{r: r, c: c}}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if grid[cur.r][cur.c] != "*" {
			grid[cur.r][cur.c] = "*"
			size++
			for _, delt := range []coord{{r: 0, c: 1}, {r: 1, c: 0}, {r: -1, c: 0}, {r: 0, c: -1}} {
				next := coord{cur.r + delt.r, cur.c + delt.c}
				if next.r < 0 || next.r > len(grid)-1 || next.c < 0 || next.c > len(grid[0])-1 || grid[next.r][next.c] == "*" || grid[next.r][next.c] == "9" {
					continue
				}
				stack = append(stack, next)
			}
		}
	}
	return size
}
