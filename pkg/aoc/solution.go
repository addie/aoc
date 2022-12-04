package aoc

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Solution struct {
	year, day int
}

func (s Solution) dataFilename() string {
	return "data/day" + strconv.Itoa(s.day) + ".txt"
}

func (s Solution) Execute() (int, int) {
	name := fmt.Sprintf("Year%dDay%d", s.year, s.day)
	inputs := []reflect.Value{reflect.ValueOf(s.dataFilename())}
	res := reflect.ValueOf(s).MethodByName(name).Call(inputs)
	return res[0].Interface().(int), res[1].Interface().(int)
}

func (s Solution) Year2022Day1(input string) (int, int) {
	strArr := readFile(input)

	currentTotal := 0
	maxTotal := 0
	var allTotals []int
	for _, line := range strArr {
		if line == "" {
			allTotals = append(allTotals, currentTotal)
			if currentTotal > maxTotal {
				maxTotal = currentTotal
			}
			currentTotal = 0
		}
		weight, _ := strconv.Atoi(line)
		currentTotal += weight
	}
	sort.Ints(allTotals)
	top3 := allTotals[len(allTotals)-1] + allTotals[len(allTotals)-2] + allTotals[len(allTotals)-3]
	return maxTotal, top3
}

func (s Solution) Year2022Day2(input string) (int, int) {
	strArr := readFile(input)

	// Opp A for Rock, B for Paper, and C for Scissors
	// Me X for Rock, Y for Paper, and Z for Scissors
	// Score 6 win, 3 tie, 0 loss
	// Score 1 for Rock, 2 for Paper, and 3 for Scissors
	part1Total := 0
	for _, line := range strArr {
		moves := strings.Split(line, " ")
		oppMove, myMove := moves[0], moves[1]
		switch myMove {
		case "X":
			part1Total += 1
		case "Y":
			part1Total += 2
		case "Z":
			part1Total += 3
		}
		if myMove == "X" && oppMove == "A" || myMove == "Y" && oppMove == "B" || myMove == "Z" && oppMove == "C" {
			part1Total += 3
		} else if myMove == "X" && oppMove == "C" || myMove == "Y" && oppMove == "A" || myMove == "Z" && oppMove == "B" {
			part1Total += 6
		}
	}

	// Part 2
	// X means you need to lose
	// Y means you need to draw
	// Z means you need to win
	// Opp A for Rock, B for Paper, and C for Scissors
	// Score 6 win, 3 tie, 0 loss
	// Score 1 for Rock, 2 for Paper, and 3 for Scissors
	part2Total := 0
	for _, line := range strArr {
		moves := strings.Split(line, " ")
		oppMove, outcome := moves[0], moves[1]
		switch outcome {
		case "X": // lose
			switch oppMove {
			case "A": // rock
				part2Total += 3
			case "B": // paper
				part2Total += 1
			case "C": // scissors
				part2Total += 2
			}
		case "Y": // draw
			switch oppMove {
			case "A": // rock
				part2Total += 4
			case "B": // paper
				part2Total += 5
			case "C": // scissors
				part2Total += 6
			}
		case "Z": // win
			switch oppMove {
			case "A": // rock
				part2Total += 8
			case "B": // paper
				part2Total += 9
			case "C": // scissors
				part2Total += 7
			}
		}
	}

	return part1Total, part2Total
}

func (s Solution) Year2022Day3(input string) (int, int) {
	strArr := readFile(input)

	for _, line := range strArr {

	}
}

func readFile(filename string) []string {
	var lines []string
	readFile, err := os.Open(filename)
	check(err)
	defer func(readFile *os.File) {
		err := readFile.Close()
		check(err)
	}(readFile)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
