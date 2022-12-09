package aoc

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day4Part1() int {
	moves, boards, err := get4()
	if err != nil {
		log.Fatal(err)
	}

	lastMove, winner := findBingoWinner(moves, boards)
	sum := sumRemainingTiles(winner)
	return sum * lastMove
}

func Day4Part2() int {
	moves, boards, err := get4()
	if err != nil {
		log.Fatal(err)
	}
	lastMove, loser := findBingoLoser(moves, boards)
	sum := sumRemainingTiles(loser)
	return sum * lastMove
}

func sumRemainingTiles(board [][]string) int {
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

func findBingoWinner(moves []string, boards [][][]string) (int, [][]string) {
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

func findBingoLoser(moves []string, boards [][][]string) (int, [][]string) {
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

func bingo(board [][]string) bool {
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

func mark(move string, board [][]string) {
	for r := range board {
		for c := range board[0] {
			if move == board[r][c] {
				board[r][c] = ""
				break
			}
		}
	}
}

func get4() ([]string, [][][]string, error) {
	file, err := os.Open("data/day4")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	moves := strings.Split(scanner.Text(), ",")
	var boards [][][]string
	var board [][]string
	for scanner.Scan() {
		row := scanner.Text()
		if row == "" {
			if len(board) > 0 {
				boards = append(boards, board)
			}
			board = [][]string{}
			continue
		}
		cellsInRow := strings.Split(row, " ")
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
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	return moves, boards, nil
}
