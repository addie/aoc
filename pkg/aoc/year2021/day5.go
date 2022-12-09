package aoc

import (
	"aoc2021/data/part"
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	r, c int
}

type coordPair struct {
	start, end coord
}

type dir int

const (
	diagonal dir = iota
	horizontal
	vertical
)

func Day5Part1() int {
	data, err := get5()
	if err != nil {
		log.Fatal(err)
	}

	maxR, maxC := getBounds(data)
	grid := createGrid(maxR, maxC)
	for _, coordPair := range data {
		plotLine(part.One, grid, coordPair)
	}
	return countIntersections(grid)
}

func Day5Part2() int {
	data, err := get5()
	if err != nil {
		log.Fatal(err)
	}
	maxR, maxC := getBounds(data)
	grid := createGrid(maxR, maxC)
	for _, coordPair := range data {
		plotLine(part.Two, grid, coordPair)
	}
	return countIntersections(grid)
}

func countIntersections(grid [][]int) int {
	numIntersections := 0
	for r := range grid {
		for c := range grid[0] {
			if grid[r][c] > 1 {
				numIntersections++
			}
		}
	}
	return numIntersections
}

func plotLine(p part.Part, grid [][]int, pair coordPair) {
	dir := getDirection(pair)
	switch dir {
	case horizontal:
		plotHorizontal(grid, pair)
	case vertical:
		plotVertical(grid, pair)
	case diagonal:
		if p == part.Two {
			plotDiagonal(grid, pair)
		}
	}
}

func plotVertical(grid [][]int, pair coordPair) {
	cur, end := pair.start.r, pair.end.r
	col := pair.start.c
	if pair.end.r < pair.start.r {
		cur, end = end, cur
	}
	for cur <= end {
		grid[cur][col] += 1
		cur++
	}
}

func plotHorizontal(grid [][]int, pair coordPair) {
	cur, end := pair.start.c, pair.end.c
	row := pair.start.r
	if pair.end.c < pair.start.c {
		cur, end = end, cur
	}
	for cur <= end {
		grid[row][cur] += 1
		cur++
	}
}

func plotDiagonal(grid [][]int, pair coordPair) {
	curCol, endCol := pair.start.c, pair.end.c
	curRow, endRow := pair.start.r, pair.end.r
	if pair.end.c < pair.start.c {
		curCol, endCol = endCol, curCol
		curRow, endRow = endRow, curRow
	}
	sub := false
	if curRow > endRow {
		sub = true
	}
	for curCol <= endCol {
		grid[curRow][curCol] += 1
		curCol++
		if sub {
			curRow--
		} else {
			curRow++
		}
	}
}

func getDirection(pair coordPair) dir {
	if pair.start.r-pair.end.r == 0 {
		return horizontal
	}
	if pair.start.c-pair.end.c == 0 {
		return vertical
	}
	return diagonal
}

func createGrid(maxR int, maxC int) [][]int {
	grid := make([][]int, maxR+1)
	for r := range grid {
		grid[r] = make([]int, maxC+1)
	}
	return grid
}

func getBounds(input []coordPair) (int, int) {
	maxR, maxC := 0, 0
	for _, coordPair := range input {
		if coordPair.start.r > maxR {
			maxR = coordPair.start.r
		}
		if coordPair.end.r > maxR {
			maxR = coordPair.end.r
		}
		if coordPair.start.c > maxC {
			maxC = coordPair.start.c
		}
		if coordPair.end.c > maxC {
			maxC = coordPair.end.c
		}
	}
	return maxR, maxC
}

func get5() ([]coordPair, error) {
	file, err := os.Open("data/day5")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var coords []coordPair
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parsedRow := strings.Split(row, "->")
		r, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parsedRow[0], ",")[0]))
		c, _ := strconv.Atoi(strings.TrimSpace(strings.Split(parsedRow[0], ",")[1]))
		start := coord{
			r: r, c: c,
		}
		r, _ = strconv.Atoi(strings.TrimSpace(strings.Split(parsedRow[1], ",")[0]))
		c, _ = strconv.Atoi(strings.TrimSpace(strings.Split(parsedRow[1], ",")[1]))
		end := coord{
			r: r, c: c,
		}
		coords = append(coords, coordPair{start, end})
	}
	return coords, nil
}
