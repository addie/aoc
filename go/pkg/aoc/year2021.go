package aoc

import (
	"log"
	"strconv"
	"strings"
)

func (s Solution[T]) Year2021Day1(input string) (int, int) {
	lines := ReadFileToLines(input)
	numIncreases1 := 0
	for i := range lines {
		if i == 0 {
			continue
		}
		if Must(strconv.Atoi(lines[i])) > Must(strconv.Atoi(lines[i-1])) {
			numIncreases1++
		}
	}

	numIncreases2 := 0
	var window []int
	runningSum := 0
	for _, line := range lines {
		lintInt := Must(strconv.Atoi(line))
		if len(window) < 3 {
			window = append(window, lintInt)
			runningSum += lintInt
			continue
		}
		nextSum := runningSum - window[0] + lintInt
		window = append(window[1:], lintInt)
		if nextSum > runningSum {
			numIncreases2++
		}
		runningSum = nextSum
	}
	return numIncreases1, numIncreases2
}

type pos struct {
	horizontal, depth, aim int
}

func (s Solution[T]) Year2021Day2(input string) (int, int) {
	lines := ReadFileToLines(input)

	res := pos{}
	for _, line := range lines {
		fullCmd := strings.Split(line, " ")
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res.horizontal += num
		case "down":
			res.depth += num
		case "up":
			res.depth -= num
		}
	}

	res2 := pos{}
	for _, line := range lines {
		fullCmd := strings.Split(line, " ")
		dir := fullCmd[0]
		num, _ := strconv.Atoi(fullCmd[1])
		switch dir {
		case "forward":
			res2.horizontal += num
			res2.depth += res2.aim * num
		case "down":
			res2.aim += num
		case "up":
			res2.aim -= num
		}
	}
	return res.horizontal * res.depth, res2.horizontal * res2.depth
}

func (s Solution[T]) Year2021Day3(input string) (int, int) {
	const (
		totalStrings = 1000
		binLength    = 12
	)

	initCount := func(count map[int]int) {
		for i := 1; i < 13; i++ {
			count[i] = 0
		}
	}
	oneCount := make(map[int]int)
	initCount(oneCount)
	lines := ReadFileToLines(input)
	for _, dat := range lines {
		for i, digit := range dat {
			if string(digit) == "1" {
				oneCount[i] += 1
			}
		}
	}
	mostCommon := ""
	leastCommon := ""
	for i := 0; i < binLength; i++ {
		if oneCount[i] > totalStrings/2 {
			mostCommon += "1"
			leastCommon += "0"
		} else {
			mostCommon += "0"
			leastCommon += "1"
		}
	}
	gamma, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	cull := func(i int, currentList []string, invariant bool, digitToKeep int) []string {
		var newList []string
		for _, binStr := range currentList {
			if invariant && string(binStr[i]) == strconv.Itoa(digitToKeep) {
				newList = append(newList, binStr)
			} else if !invariant && string(binStr[i]) == strconv.Itoa(1-digitToKeep) {
				newList = append(newList, binStr)
			}
		}
		return newList
	}
	buildOxygenNum := func(currentList []string) int {
		oneCount := 0
		var oxygen int64
		for i := 0; i < binLength; i++ {
			for _, binStr := range currentList {
				if string(binStr[i]) == "1" {
					oneCount += 1
				}
			}
			keepOnes := oneCount >= len(currentList)-oneCount
			oneCount = 0
			currentList = cull(i, currentList, keepOnes, 1)
			if len(currentList) == 1 {
				var err error
				oxygen, err = strconv.ParseInt(currentList[0], 2, 64)
				if err != nil {
					log.Fatal(err)
				}
				return int(oxygen)
			}
		}
		return -1
	}

	buildCO2Num := func(currentList []string) int {
		zeroCount := 0
		var co2 int64
		for i := 0; i < binLength; i++ {
			for _, binStr := range currentList {
				if string(binStr[i]) == "0" {
					zeroCount += 1
				}
			}
			keepZeros := zeroCount <= len(currentList)-zeroCount
			zeroCount = 0
			currentList = cull(i, currentList, keepZeros, 0)
			if len(currentList) == 1 {
				var err error
				co2, err = strconv.ParseInt(currentList[0], 2, 64)
				if err != nil {
					log.Fatal(err)
				}
				return int(co2)
			}
		}
		return -1
	}
	oxygen := buildOxygenNum(lines)
	co2 := buildCO2Num(lines)
	return int(gamma) * int(epsilon), oxygen * co2
}

func (s Solution[T]) Year2021Day4(input string) (int, int) {
	lines := ReadFileToLines(input)
	var moves []string
	var boards [][][]string
	var board [][]string
	for i, line := range lines {
		if i == 0 {
			moves = strings.Split(line, ",")
			continue
		}
		if line == "" {
			if len(board) > 0 {
				boards = append(boards, board)
			}
			board = [][]string{}
			continue
		}
		cellsInRow := strings.Split(line, " ")
		compress := func(s []string) []string {
			var r []string
			for _, str := range s {
				if str != "" {
					r = append(r, str)
				}
			}
			return r
		}
		cellsInRow = compress(cellsInRow)
		board = append(board, cellsInRow)
	}
	boards = append(boards, board)

	sumRemainingTiles := func(board [][]string) int {
		sum := 0
		for r := range board {
			for c := range board[0] {
				if board[r][c] != "" {
					val, _ := strconv.Atoi(board[r][c])
					sum += val
				}
			}
		}
		return sum
	}

	bingo := func(board [][]string) bool {
		for r := range board {
			for c := range board[0] {
				if board[r][c] != "" {
					break
				}
				if c == len(board[0])-1 {
					return true
				}
			}
		}
		for c := range board[0] {
			for r := range board {
				if board[r][c] != "" {
					break
				}
				if r == len(board[0])-1 {
					return true
				}
			}
		}
		return false
	}

	mark := func(move string, board [][]string) {
		for r := range board {
			for c := range board[0] {
				if move == board[r][c] {
					board[r][c] = ""
					break
				}
			}
		}
	}

	findBingoWinner := func(moves []string, boards [][][]string) (int, [][]string) {
		for _, move := range moves {
			for _, board := range boards {
				mark(move, board)
				if bingo(board) {
					lastMove, _ := strconv.Atoi(move)
					return lastMove, board
				}
			}
		}
		log.Fatal("no winner")
		return 0, nil
	}

	findBingoLoser := func(moves []string, boards [][][]string) (int, [][]string) {
		winners := make(map[int]bool)
		for _, move := range moves {
			for i, board := range boards {
				if winners[i] {
					continue
				}
				mark(move, board)
				if bingo(board) {
					winners[i] = true
					if len(winners) == len(boards) {
						lastMove, _ := strconv.Atoi(move)
						return lastMove, boards[i]
					}
				}
			}
		}
		log.Fatal("no loser")
		return 0, nil
	}

	lastMoveWinner, winner := findBingoWinner(moves, boards)
	sumWinner := sumRemainingTiles(winner)
	lastMoveLoser, loser := findBingoLoser(moves, boards)
	sumLoser := sumRemainingTiles(loser)
	return sumWinner * lastMoveWinner, sumLoser * lastMoveLoser
}

func (s Solution[T]) Year2021Day5(input string) (int, int) {
	type coord struct{ r, c int }
	type coordPair struct{ start, end coord }

	const (
		diagonal   string = "diagonal"
		horizontal string = "horizontal"
		vertical   string = "vertical"
	)
	createGrid := func(maxR int, maxC int) [][]int {
		grid := make([][]int, maxR+1)
		for r := range grid {
			grid[r] = make([]int, maxC+1)
		}
		return grid
	}

	getBounds := func(input []coordPair) (int, int) {
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

	countIntersections := func(grid [][]int) int {
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

	plotVertical := func(grid [][]int, pair coordPair) {
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

	plotHorizontal := func(grid [][]int, pair coordPair) {
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

	plotDiagonal := func(grid [][]int, pair coordPair) {
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
	getDirection := func(pair coordPair) string {
		if pair.start.r-pair.end.r == 0 {
			return horizontal
		}
		if pair.start.c-pair.end.c == 0 {
			return vertical
		}
		return diagonal
	}
	plotLine := func(part int, grid [][]int, pair coordPair) {
		dir := getDirection(pair)
		switch dir {
		case horizontal:
			plotHorizontal(grid, pair)
		case vertical:
			plotVertical(grid, pair)
		case diagonal:
			if part == 2 {
				plotDiagonal(grid, pair)
			}
		}
	}

	var cp []coordPair
	lines := ReadFileToLines(input)
	for _, line := range lines {
		row := strings.Split(line, " -> ")
		cp = append(cp, coordPair{
			start: coord{
				r: Must(strconv.Atoi(strings.Split(row[0], ",")[0])),
				c: Must(strconv.Atoi(strings.Split(row[0], ",")[1])),
			},
			end: coord{
				r: Must(strconv.Atoi(strings.Split(row[1], ",")[0])),
				c: Must(strconv.Atoi(strings.Split(row[1], ",")[1])),
			},
		})
	}
	maxR, maxC := getBounds(cp)
	grid := createGrid(maxR, maxC)
	for _, coordPair := range cp {
		plotLine(1, grid, coordPair)
	}
	maxR2, maxC2 := getBounds(cp)
	grid2 := createGrid(maxR2, maxC2)
	for _, coordPair := range cp {
		plotLine(2, grid2, coordPair)
	}
	return countIntersections(grid), countIntersections(grid2)
}

//
// func (s Solution[T]) Year2021Day6(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day7(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day8(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day9(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day10(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day11(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day12(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day13(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day14(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day15(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day16(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
//
// func (s Solution[T]) Year2021Day17(input string) (int, int) {
// 	lines := ReadFileToLines(input)
//
// }
